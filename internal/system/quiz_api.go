package system

import (
	"quiz-system/internal/application"
	concurrentmap "quiz-system/internal/concurrent_map"
	"quiz-system/internal/rest"
	"quiz-system/internal/rest/quiz"
)

type QuizApi struct {
	system *System
	api    *rest.Api
	app    application.App
}

func NewQuizApi(system *System) *QuizApi {

	mux := system.Mux()

	questionRepository := concurrentmap.NewQuestionRepository()
	evaluationRepository := concurrentmap.NewEvaluationRepository(system.logger)

	application := application.NewApplication(questionRepository, evaluationRepository)
	quizV1 := quiz.NewQuizV1(application)
	api := rest.NewApi(mux, quizV1)

	return &QuizApi{
		system: system,
		api:    api,
		app:    application,
	}
}

func (q *QuizApi) Api() *rest.Api {
	return q.api
}

func (q *QuizApi) Init() {
	q.api.Init()
}
