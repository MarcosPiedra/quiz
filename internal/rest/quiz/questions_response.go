package quiz

import (
	"net/http"
	"quiz-system/internal/domain"

	"github.com/go-chi/render"
)

type QuestionsResponse struct {
	Questions []QuestionResponse `json:"questions"`
}

type QuestionResponse struct {
	Id             int               `json:"id"`
	Question       string            `json:"question"`
	PosibleAnswers []AnswersResponse `json:"posibleAnswers"`
}

type AnswersResponse struct {
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

func NewQuestionsResponse(questions []*domain.Question) QuestionsResponse {
	response := QuestionsResponse{}
	response.Questions = make([]QuestionResponse, len(questions))
	for idx, q := range questions {
		response.Questions[idx] = NewQuestionResponse(q)
	}
	return response
}

func NewQuestionResponse(question *domain.Question) QuestionResponse {
	response := QuestionResponse{
		Id:       question.Id,
		Question: question.Question,
	}
	response.PosibleAnswers = make([]AnswersResponse, len(question.PossibleAnswers))
	for idx, a := range question.PossibleAnswers {
		response.PosibleAnswers[idx].Answer = a.Answer
		response.PosibleAnswers[idx].Id = a.Id
	}

	return response
}

func (QuestionsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}

func (QuestionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}
