package llm

type LLMClient interface {
	GenerateCommitMessage(diff string) (string, error)
}
