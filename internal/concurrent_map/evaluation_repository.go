package concurrentmap

import (
	"context"
	"quiz-system/internal/domain"

	"github.com/rs/zerolog"
)

type EvaluationRepository struct {
	cm     *ConcurrentMap[domain.Evaluation]
	logger zerolog.Logger
}

// NextId implements domain.EvaluationRepository.
func (e *EvaluationRepository) NextId() int {
	return e.cm.NextId()
}

func (e *EvaluationRepository) Add(evaluation *domain.Evaluation, ctx context.Context) error {
	return e.cm.Set(evaluation.Id, *evaluation, ctx)
}

func (e *EvaluationRepository) GetAll(ctx context.Context) ([]*domain.Evaluation, error) {
	evals, err := e.cm.GetAll(ctx)
	var pointers []*domain.Evaluation
	for i := 0; i < len(evals); i++ {
		pointers = append(pointers, &evals[i])
	}
	return pointers, err
}

func NewEvaluationRepository(logger zerolog.Logger) *EvaluationRepository {
	return &EvaluationRepository{
		cm:     NewConcurrentMap[domain.Evaluation](),
		logger: logger,
	}
}

var _ domain.EvaluationRepository = (*EvaluationRepository)(nil)
