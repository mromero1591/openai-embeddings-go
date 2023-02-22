package openaiembeddingsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	APIKey         string               `json:"apiKey"`
	EmbeddingModel OpenAIEmbeddingModel `json:"embeddingModel"`
}

const (
	ADA002                   OpenAIEmbeddingModel = "text-embedding-ada-002"
	createEmbeddingsEndPoint string               = "https://api.openai.com/v1/embeddings"
)

func NewClient(apiKey string) Client {
	return Client{
		APIKey: apiKey,
	}
}

func (c *Client) WithModel(model OpenAIEmbeddingModel) {
	c.EmbeddingModel = model
}

func (c Client) CreateOpenAIEmbeddings(input string) (OpenAIEmbeddingResponse, error) {
	requestBody := OpenAIEmbeddingRequest{
		Model: c.EmbeddingModel,
		Input: input,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest(http.MethodPost, createEmbeddingsEndPoint, bytes.NewReader(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return OpenAIEmbeddingResponse{}, err
	}
	defer response.Body.Close()

	var openaiResponse OpenAIEmbeddingResponse
	if err := json.NewDecoder(response.Body).Decode(&openaiResponse); err != nil {
		return OpenAIEmbeddingResponse{}, err
	}

	return openaiResponse, nil
}
