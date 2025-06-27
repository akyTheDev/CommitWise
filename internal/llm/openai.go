package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type openAiRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAiRequest struct {
	Model       string                 `json:"model"`
	Messages    []openAiRequestMessage `json:"messages"`
	Temperature float64                `json:"temperature"`
}

type openAiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		}
	} `json:"choices"`
	Usage struct {
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}

const (
	openAIAPIURL   = "https://api.openai.com/v1/chat/completions"
	defaultModel   = "gpt-4o"
	defaultTimeout = 30 * time.Second
)

type OpenAiClient struct {
	token      string
	httpClient *http.Client
}

func NewOpenAiClient(token string) *OpenAiClient {
	return &OpenAiClient{
		token: token,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (client *OpenAiClient) GenerateCommitMessage(diff string) (string, error) {
	prompt := fmt.Sprintf(`
	You are an expert at writing commit messages that follow the Conventional Commits specification.
	Your task is to analyze the following git diff and generate a complete and high-quality commit message.

	Follow these rules strictly:
	1. The final output must be only the raw commit message, with no extra text, explanations, or markdown formatting.
	2. The message must adhere to the Conventional Commits format: <type>[optional scope]: <description>\n\n[optional body]\n\n[optional footer(s)].
	3. The <type> must be one of: feat, fix, build, chore, ci, docs, style, refactor, perf, test.
	4. The <description> must be a concise summary, written in the imperative mood (e.g., "add login page" not "added login page").
	5. If the changes are complex, non-trivial, or require motivation, provide a detailed [optional body] to explain the 'what' and 'why'.
	6. If the diff introduces a backward-incompatible change, you MUST include a footer in the format 'BREAKING CHANGE: <description of the breaking change>'.
	7. Only include the optional scope, body, and footer if they are necessary based on the provided diff.

	Here is the git diff:
	---
	%s
	---`, diff)

	reqPayload := openAiRequest{
		Model: "gpt-4o",
		Messages: []openAiRequestMessage{
			{
				Role:    "system",
				Content: "You are an expert at writing commit messages based on git diffs",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.2,
	}

	reqBody, err := json.Marshal(reqPayload)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", openAIAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("Failed to create http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.token)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Api request failed with status code: %d", resp.StatusCode)
	}

	var response openAiResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("Failed to decode response: %w", err)
	}

	if len(response.Choices) == 0 || response.Choices[0].Message.Content == "" {
		return "", errors.New("No commit message content returned from api")
	}

	return response.Choices[0].Message.Content, nil
}
