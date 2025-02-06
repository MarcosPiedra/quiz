package domain

import (
	"context"
)

type QuestionRepository interface {
	GetAll(ctx context.Context) ([]*Question, error)
	GetById(id int, ctx context.Context) (*Question, error)
	Count(ctx context.Context) (int, error)
}
