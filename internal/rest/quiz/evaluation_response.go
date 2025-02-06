package quiz

import (
	"fmt"
	"net/http"
	"quiz-system/internal/domain"

	"github.com/go-chi/render"
)

type EvaluationResponse struct {
	Id                int    `json:"id"`
	TotalQuestions    int    `json:"totalQuestions"`
	QuestionsAnswered int    `json:"questionsAnswered"`
	CorrectAnswer     int    `json:"correctAnswer"`
	Comparative       string `json:"comparative"`
}

func NewEvaluationResponse(evaluation *domain.Evaluation) EvaluationResponse {
	comparative := ""
	switch evaluation.PercentBetterThanOthers {
	case -1:
		comparative = "You are the first to be quizzed!"
	case 0:
		comparative = "Please, try again ;)"
	default:
		comparative = fmt.Sprintf("You were better than %d of all quizzers", evaluation.PercentBetterThanOthers)
	}
	return EvaluationResponse{
		Id:                evaluation.Id,
		TotalQuestions:    evaluation.TotalQuestions,
		QuestionsAnswered: evaluation.QuestionsAnswered,
		CorrectAnswer:     evaluation.CorrectAnswer,
		Comparative:       comparative,
	}
}

func (EvaluationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}
