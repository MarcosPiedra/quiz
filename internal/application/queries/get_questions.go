package queries

import (
	"context"
	"quiz-system/internal/domain"
)

type (
	GetQuestionsHandler struct {
		questionsRepository domain.QuestionRepository
	}
)

func NewGetQuestionsHandler(questionsRepository domain.QuestionRepository) GetQuestionsHandler {
	return GetQuestionsHandler{questionsRepository: questionsRepository}
}

func (h GetQuestionsHandler) GetQuestions(ctx context.Context) ([]*domain.Question, error) {
	return h.questionsRepository.GetAll(ctx)
}
