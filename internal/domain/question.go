package domain

type Question struct {
	Id              int
	Question        string
	PossibleAnswers []PossibleAnswer
}

type PossibleAnswer struct {
	Id        int
	Answer    string
	IsCorrect bool
}

func (q Question) ExistAswerId(answerId int) bool {
	for i := 0; i < len(q.PossibleAnswers); i++ {
		if q.PossibleAnswers[i].Id == answerId {
			return true
		}
	}
	return false
}

func (q Question) IsCorrectAnswer(answerId int) bool {
	for i := 0; i < len(q.PossibleAnswers); i++ {
		if q.PossibleAnswers[i].Id == answerId {
			return q.PossibleAnswers[i].IsCorrect
		}
	}
	return false
}
