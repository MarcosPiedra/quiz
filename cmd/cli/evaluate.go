package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	q "quiz-system/internal/rest/quiz"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Evaluate your answers, please use this format '1 3 2 4'",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		param := args[0]

		ids := strings.Fields(param)
		if len(ids) == 0 {
			fmt.Println("Error: The argument must be a pattern like '1 2 3 4'")
			return
		}

		evaluations := make([]q.EvaluationRequest, len(ids))
		for idx, id := range ids {
			aId, err := strconv.Atoi(id)
			if err != nil || aId < 0 {
				fmt.Println("Error: The argument must be a positive integer")
				return
			}
			evaluations[idx].QuestionId = idx + 1
			evaluations[idx].AnswerId = aId
		}
		payload := q.EvaluationsRequest{Answers: evaluations}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		resp, err := http.Post("http://localhost:8080/quiz/v1/evaluation", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error sending POST request:", err)
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

		var response q.EvaluationResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Printf("TotalQuestions: %d\n", response.TotalQuestions)
		fmt.Printf("QuestionsAnswered: %d\n", response.QuestionsAnswered)
		fmt.Printf("CorrectAnswer: %d\n", response.CorrectAnswer)
		fmt.Printf("Comparative: %s\n", response.Comparative)
	},
}

func init() {
	RootCmd.AddCommand(evaluateCmd)
}
