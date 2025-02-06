package commands

import (
	"context"
	"quiz-system/internal/domain"

	err "github.com/stackus/errors"
)

type (
	AddEvaluation struct {
		Answers []Answers
	}

	Answers struct {
		QuestionId int
		AnswerId   int
	}

	AddEvaluationHandler struct {
		evaluationRepository domain.EvaluationRepository
		questionRepository   domain.QuestionRepository
	}
)

func NewAddEvaluationHandler(evaluationRepository domain.EvaluationRepository, questionRepository domain.QuestionRepository) AddEvaluationHandler {
	return AddEvaluationHandler{evaluationRepository: evaluationRepository, questionRepository: questionRepository}
}

func (h AddEvaluationHandler) AddEvaluation(cmd AddEvaluation, ctx context.Context) (*domain.Evaluation, error) {

	var totalQuestions = 0
	var questionsAnswered = 0
	var correctAnswer = 0

	totalQuestions, _ = h.questionRepository.Count(ctx)

	for i := 0; i < len(cmd.Answers); i++ {
		a := cmd.Answers[i]
		q, e := h.questionRepository.GetById(a.QuestionId, ctx)
		if e != nil {
			return nil, err.Wrap(err.ErrBadRequest, "Invalid question")
		}

		if !q.ExistAswerId(a.AnswerId) {
			return nil, err.Wrap(err.ErrBadRequest, "Invalid answer")
		}

		questionsAnswered++

		if q.IsCorrectAnswer(a.AnswerId) {
			correctAnswer++
		}

	}

	others, _ := h.evaluationRepository.GetAll(ctx)
	scores := make([]float32, len(others))
	for _, o := range others {
		scores = append(scores, o.Score)
	}

	evaluationToAdd := domain.NewEvaluation(h.evaluationRepository.NextId(), totalQuestions, questionsAnswered, correctAnswer, scores)

	e := h.evaluationRepository.Add(evaluationToAdd, ctx)

	if e != nil {
		return nil, err.Wrap(err.ErrInternalServerError, e.Error())
	}

	return evaluationToAdd, nil
}
