package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/akyTheDev/CommitWise/internal/config"
	"github.com/akyTheDev/CommitWise/internal/llm"
)

func main() {
	if !isPiped() {
		fmt.Println("No data piped. This command expects input from a pipe.")
		fmt.Println("Example: git diff | commitwise")
		os.Exit(1)
	}

	pipedInput, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read from pipe: %v", err)
	}

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	llmClient := llm.NewOpenAiClient(config.OPENAI_API_KEY)

	commitMessage, err := llmClient.GenerateCommitMessage(string(pipedInput))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commitMessage)
}

// isPiped checks if the application is receiving data from a pipe.
func isPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}
