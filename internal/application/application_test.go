package application_test

import (
	"context"
	"quiz-system/internal/application"
	"quiz-system/internal/application/commands"
	"quiz-system/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para QuestionRepository
type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) GetAll(ctx context.Context) ([]*domain.Question, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Question), args.Error(1)
}

func (m *MockQuestionRepository) GetById(id int, ctx context.Context) (*domain.Question, error) {
	args := m.Called(id, ctx)
	return args.Get(0).(*domain.Question), args.Error(1)
}

func (m *MockQuestionRepository) Count(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

// Mock para EvaluationRepository
type MockEvaluationRepository struct {
	mock.Mock
}

func (m *MockEvaluationRepository) GetAll(ctx context.Context) ([]*domain.Evaluation, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Evaluation), args.Error(1)
}

func (m *MockEvaluationRepository) Add(evaluation *domain.Evaluation, ctx context.Context) error {
	args := m.Called(evaluation, ctx)
	return args.Error(0)
}

func (m *MockEvaluationRepository) NextId() int {
	args := m.Called()
	return args.Int(0)
}

func TestGetQuestions(t *testing.T) {
	mockQuestionRepo := new(MockQuestionRepository)
	app := application.NewApplication(mockQuestionRepo, new(MockEvaluationRepository))

	ctx := context.Background()
	mockQuestions := make([]*domain.Question, 1)
	mockQuestions = append(mockQuestions, &domain.Question{
		Id:              1,
		Question:        "Question 1",
		PossibleAnswers: []domain.PossibleAnswer{},
	})
	mockQuestionRepo.On("GetAll", ctx).Return(mockQuestions, nil)

	questions, err := app.GetQuestions(ctx)
	assert.NoError(t, err)
	assert.Equal(t, mockQuestions, questions)

	mockQuestionRepo.AssertExpectations(t)
}

func TestGetQuestion(t *testing.T) {
	mockQuestionRepo := new(MockQuestionRepository)
	app := application.NewApplication(mockQuestionRepo, new(MockEvaluationRepository))

	ctx := context.Background()
	mockQuestion := &domain.Question{
		Id:              1,
		Question:        "Question 1",
		PossibleAnswers: []domain.PossibleAnswer{},
	}
	mockQuestionRepo.On("GetById", 1, ctx).Return(mockQuestion, nil)

	question, err := app.GetQuestion(1, ctx)
	assert.NoError(t, err)
	assert.Equal(t, mockQuestion, question)

	mockQuestionRepo.AssertExpectations(t)
}

func TestAddEvaluation(t *testing.T) {
	mockEvaluationRepo := new(MockEvaluationRepository)
	mockQuestionRepo := new(MockQuestionRepository)
	app := application.NewApplication(mockQuestionRepo, mockEvaluationRepo)

	ctx := context.Background()
	mockEvaluationToAdd := &domain.Evaluation{Id: 2, TotalQuestions: 2, QuestionsAnswered: 2, CorrectAnswer: 1, Score: 0.5, PercentBetterThanOthers: 50}
	mockEvaluation := &domain.Evaluation{Id: 1, TotalQuestions: 2, QuestionsAnswered: 2, CorrectAnswer: 1, Score: 0.5, PercentBetterThanOthers: 50}
	mockEvaluations := []*domain.Evaluation{mockEvaluation}
	mockEvaluationRepo.On("Add", mockEvaluationToAdd, ctx).Return(nil)
	mockEvaluationRepo.On("GetAll", ctx).Return(mockEvaluations, nil)
	mockEvaluationRepo.On("NextId").Return(2)

	question1 := &domain.Question{
		Id:       1,
		Question: "Question 1",
		PossibleAnswers: []domain.PossibleAnswer{
			{
				Id:        1,
				Answer:    "Answer 1 1",
				IsCorrect: true,
			},
			{
				Id:        2,
				Answer:    "Answer 1 2",
				IsCorrect: false,
			},
		},
	}
	question2 := &domain.Question{
		Id:       2,
		Question: "Question 2",
		PossibleAnswers: []domain.PossibleAnswer{
			{
				Id:        1,
				Answer:    "Answer 2 1",
				IsCorrect: false,
			},
			{
				Id:        2,
				Answer:    "Answer 2 2",
				IsCorrect: true,
			},
		},
	}
	mockQuestions := []*domain.Question{
		question1,
		question2,
	}

	mockQuestionRepo.On("Count", ctx).Return(len(mockQuestions), nil)
	mockQuestionRepo.On("GetById", 1, ctx).Return(question1, nil)
	mockQuestionRepo.On("GetById", 2, ctx).Return(question2, nil)

	evaluation, err := app.AddEvaluation(commands.AddEvaluation{
		Answers: []commands.Answers{
			{
				QuestionId: 1,
				AnswerId:   1,
			},
			{
				QuestionId: 2,
				AnswerId:   1,
			},
		},
	}, ctx)

	assert.Equal(t, mockEvaluationToAdd, evaluation)
	assert.NoError(t, err)

	mockEvaluationRepo.AssertExpectations(t)
}
