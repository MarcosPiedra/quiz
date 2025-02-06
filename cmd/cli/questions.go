package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	q "quiz-system/internal/rest/quiz"

	"github.com/spf13/cobra"
)

var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "Get all questions",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/quiz/v1/questions")
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()
		if resp.Status != "200 OK" {
			var responseErr q.ErrorResponse
			if err := json.NewDecoder(resp.Body).Decode(&responseErr); err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}			
			fmt.Println("Status:", responseErr.HTTPStatusCode)
			fmt.Println("Status error:", responseErr.Error)
			return
		}

		var response q.QuestionsResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		for _, question := range response.Questions {
			fmt.Printf("%d %s\n", question.Id, question.Question)
			for _, answer := range question.PosibleAnswers {
				fmt.Printf("- %d %s\n", answer.Id, answer.Answer)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(questionsCmd)
}
