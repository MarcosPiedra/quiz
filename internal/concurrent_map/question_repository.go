package concurrentmap

import (
	"context"
	"errors"
	"quiz-system/internal/domain"
)

type QuestionRepository struct {
	cm *ConcurrentMap[domain.Question]
}

func (q *QuestionRepository) InitDb() {
	ctx := context.Background()

	q.cm.Set(1, domain.Question{
		Id:       1,
		Question: "Which planet is closest to the sun?",
		PossibleAnswers: []domain.PossibleAnswer{
			newAnswer(1, "Earth", false),
			newAnswer(2, "Venus", false),
			newAnswer(3, "Mercury", true),
		},
	}, ctx)

	q.cm.Set(2, domain.Question{
		Id:       2,
		Question: "Which of the following is the largest ocean on Earth?",
		PossibleAnswers: []domain.PossibleAnswer{
			newAnswer(1, "Atlantic Ocean", false),
			newAnswer(2, "Indian Ocean", false),
			newAnswer(3, "Pacific Ocean", true),
		},
	}, ctx)

	q.cm.Set(3, domain.Question{
		Id:       3,
		Question: "Which of the following elements is a noble gas?",
		PossibleAnswers: []domain.PossibleAnswer{
			newAnswer(1, "Helium", true),
			newAnswer(2, "Oxygen", false),
			newAnswer(3, "Nitrogen", false),
		},
	}, ctx)

	q.cm.Set(4, domain.Question{
		Id:       4,
		Question: "Which famous scientist developed the theory of relativity?",
		PossibleAnswers: []domain.PossibleAnswer{
			newAnswer(1, "Isaac Newton", false),
			newAnswer(2, "Albert Einstein", true),
			newAnswer(3, "Galileo Galilei", false),
		},
	}, ctx)
}

func newAnswer(id int, answer string, isCorrect bool) domain.PossibleAnswer {
	return domain.PossibleAnswer{
		Id:        id,
		Answer:    answer,
		IsCorrect: isCorrect,
	}
}

// Count implements domain.QuestionRepository.
func (q *QuestionRepository) Count(ctx context.Context) (int, error) {
	i, e := q.cm.Count(ctx)
	return i, e
}

// GetAll implements QuestionRepository.
func (q *QuestionRepository) GetAll(ctx context.Context) ([]*domain.Question, error) {
	qs, err := q.cm.GetAll(ctx)
	var pointers []*domain.Question
	for _, question := range qs {
		pointers = append(pointers, &question)
	}

	return pointers, err
}

// GetById implements QuestionRepository.
func (q *QuestionRepository) GetById(id int, ctx context.Context) (*domain.Question, error) {
	qu, exists, err := q.cm.Get(id, ctx)
	if !exists {
		return nil, errors.New("element not exists")
	}
	return &qu, err
}

func NewQuestionRepository() *QuestionRepository {
	q := &QuestionRepository{
		cm: NewConcurrentMap[domain.Question](),
	}
	q.InitDb()
	return q
}

var _ domain.QuestionRepository = (*QuestionRepository)(nil)
