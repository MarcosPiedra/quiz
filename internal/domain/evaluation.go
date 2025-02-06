package domain

type Evaluation struct {
	Id                      int
	TotalQuestions          int
	QuestionsAnswered       int
	CorrectAnswer           int
	Score                   float32
	PercentBetterThanOthers int
}

func NewEvaluation(id, totalQuestions, questionsAnswered, correctAnswer int, others []float32) *Evaluation {
	score := float32(correctAnswer) / float32(totalQuestions)

	return &Evaluation{
		Id:                      id,
		TotalQuestions:          totalQuestions,
		QuestionsAnswered:       questionsAnswered,
		CorrectAnswer:           correctAnswer,
		Score:                   score,
		PercentBetterThanOthers: CalculateComparation(others, score),
	}
}

func CalculateComparation(others []float32, score float32) int {
	if len(others) == 0 {
		return -1
	}

	count := 0
	for _, s := range others {
		if s < score {
			count++
		}
	}

	return int((float32(count) / float32(len(others))) * 100)
}
