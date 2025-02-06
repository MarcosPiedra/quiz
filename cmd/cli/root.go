package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "CLI application with Cobra",
	Long:  "Example application using Cobra for command handling.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `quiz questions` to retrieve all questions.")
		fmt.Println("Use `quiz question 1` to retrieve all answers to a concrete question.")
		fmt.Println("Use `quiz evaluate \"1 2 3 4\"` to evaluate your answer.")
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
