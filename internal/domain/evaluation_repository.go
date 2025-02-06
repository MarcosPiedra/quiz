package domain

import (
	"context"
)

type EvaluationRepository interface {
	GetAll(ctx context.Context) ([]*Evaluation, error)
	Add(evaluation *Evaluation, ctx context.Context) error
	NextId() int
}
