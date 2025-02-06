package queries

import (
	"context"
	"quiz-system/internal/domain"
)

type (
	GetQuestionHandler struct {
		questionsRepository domain.QuestionRepository
	}
)

func NewGetQuestionHandler(questionsRepository domain.QuestionRepository) GetQuestionHandler {
	return GetQuestionHandler{questionsRepository: questionsRepository}
}

func (h GetQuestionHandler) GetQuestion(id int, ctx context.Context) (*domain.Question, error) {
	return h.questionsRepository.GetById(id, ctx)
}
