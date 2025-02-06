package quiz

import (
	"quiz-system/internal/application/commands"
)

type EvaluationsRequest struct {
	Answers []EvaluationRequest `json:"answers"`
}

type EvaluationRequest struct {
	QuestionId int `json:"questionId" validate:"required"`
	AnswerId   int `json:"answerId" validate:"required"`
}

func (r EvaluationsRequest) ToCommand() commands.AddEvaluation {
	var target commands.AddEvaluation
	for _, ans := range r.Answers {
		target.Answers = append(target.Answers, commands.Answers{
			QuestionId: ans.QuestionId,
			AnswerId:   ans.AnswerId,
		})
	}
	return target
}
