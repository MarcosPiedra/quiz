package system

import (
	"net/http"
	"quiz-system/internal/config"
	"quiz-system/internal/logger"
	"quiz-system/internal/web/static"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

type System struct {
	cfg    config.AppConfig
	mux    *chi.Mux
	logger zerolog.Logger
}

func NewSystem(cfg config.AppConfig) (*System, error) {
	var system = &System{}

	system.cfg = cfg

	system.initLog()
	system.initMux()

	return system, nil
}

func (s *System) initLog() {
	s.logger = logger.NewLogger(logger.LogConfig{
		Environment: s.cfg.Environment,
		LogLevel:    logger.Level(s.cfg.LogLevel),
	})
}

func (s *System) initMux() {
	s.mux = chi.NewMux()
	s.mux.Use(middleware.URLFormat)
	s.mux.Use(render.SetContentType(render.ContentTypeJSON))
	s.mux.Use(middleware.Recoverer)
}

func (s *System) Mux() *chi.Mux {
	return s.mux
}

func (s *System) Cfg() config.AppConfig {
	return s.cfg
}

func (s *System) Logger() zerolog.Logger {
	return s.logger
}

func (s *System) StartWebServer() {
	const swaggerPath = "/swagger/"
	s.mux.Mount(swaggerPath, http.StripPrefix(swaggerPath, http.FileServer(http.FS(static.SwaggerIndex))))
	const swaggerUiPath = "/swagger-ui/"
	s.mux.Mount(swaggerUiPath, http.FileServer(http.FS(static.SwaggerUi)))

	webServer := &http.Server{
		Addr:    s.cfg.Web.Address(),
		Handler: s.mux,
	}

	s.logger.Info().Msgf("** web server started; listening at http://localhost%s\n", s.cfg.Web.Port)
	defer s.logger.Info().Msg("** web server shutdown")
	if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Err(err)

		return
	}
}

func (s *System) Shutdown() {
}
