package rest

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Controller interface {
		Register(router chi.Router)
	}
	QuizV1 interface {
		Controller
	}

	Api struct {
		mux    *chi.Mux
		QuizV1 QuizV1
	}
)

//go:embed swagger.json
var swagger embed.FS

// @title           Quiz
// @version         1.0

// @contact.name   Marcos
// @contact.email  piedra.osuna@gmail.com

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func NewApi(
	mux *chi.Mux,
	quizV1 QuizV1,
) *Api {

	return &Api{
		mux:    mux,
		QuizV1: quizV1,
	}
}

func (api *Api) Init() {
	const specRoot = "/quiz"

	// API version 1
	api.mux.Route("/quiz/v1", func(r chi.Router) {
		api.QuizV1.Register(r)
	})

	api.mux.Mount(specRoot, http.StripPrefix(specRoot, http.FileServer(http.FS(swagger))))
}
