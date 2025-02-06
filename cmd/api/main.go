package main

import (
	"fmt"
	"os"
	"quiz-system/internal/config"
	s "quiz-system/internal/system"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("backend exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() (err error) {

	var cfg config.AppConfig
	cfg, err = config.Setup()
	if err != nil {
		return err
	}

	sys, err := s.NewSystem(cfg)
	if err != nil {
		return err
	}

	defer func(s *s.System) {
		s.Shutdown()
	}(sys)

	quizApi := s.NewQuizApi(sys)
	quizApi.Init()

	fmt.Println("starting quiz application")
	defer fmt.Println("stopped quiz application")

	sys.StartWebServer()

	return nil
}
