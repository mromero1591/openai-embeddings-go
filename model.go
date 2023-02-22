package openaiembeddingsgo

type (
	OpenAIEmbeddingModel string

	OpenAIEmbeddingRequest struct {
		Model OpenAIEmbeddingModel `json:"model"`
		Input string               `json:"input"`
	}

	OpenAIEmbeddingResponse struct {
		Object string               `json:"object"`
		Data   []OpenAIEmbedding    `json:"data"`
		Model  OpenAIEmbeddingModel `json:"model"`
		Usage  OpenAIUsage          `json:"usage"`
	}

	OpenAIUsage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	}

	OpenAIEmbedding struct {
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	}
)
