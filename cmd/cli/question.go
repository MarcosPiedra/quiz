package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	q "quiz-system/internal/rest/quiz"

	"github.com/spf13/cobra"
)

var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "Get specific question (e.g. '1')",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil || id <= 0 {
			fmt.Println("Error: The argument must be a positive integer")
			return
		}

		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/quiz/v1/questions/%d", id))
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
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

		defer resp.Body.Close()

		var response q.QuestionResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Printf("%d %s\n", response.Id, response.Question)
		for _, answer := range response.PosibleAnswers {
			fmt.Printf("- %d %s\n", answer.Id, answer.Answer)
		}
	},
}

func init() {
	RootCmd.AddCommand(questionCmd)
}
