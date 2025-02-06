package application

import (
	"context"
	"quiz-system/internal/application/commands"
	"quiz-system/internal/application/queries"
	"quiz-system/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		AddEvaluation(cmd commands.AddEvaluation, ctx context.Context) (*domain.Evaluation, error)
	}

	Queries interface {
		GetQuestions(ctx context.Context) ([]*domain.Question, error)
		GetQuestion(id int, ctx context.Context) (*domain.Question, error)
	}

	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.AddEvaluationHandler
	}

	appQueries struct {
		queries.GetQuestionsHandler
		queries.GetQuestionHandler
	}
)

var _ App = (*Application)(nil)

func NewApplication(
	questionRespository domain.QuestionRepository,
	evaluationRepository domain.EvaluationRepository,
) *Application {
	return &Application{
		appCommands: appCommands{
			AddEvaluationHandler: commands.NewAddEvaluationHandler(evaluationRepository, questionRespository),
		},
		appQueries: appQueries{
			GetQuestionsHandler: queries.NewGetQuestionsHandler(questionRespository),
			GetQuestionHandler:  queries.NewGetQuestionHandler(questionRespository),
		},
	}
}
