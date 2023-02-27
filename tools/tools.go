package tools

import (
	"fmt"
	"math"
	"strings"

	openaiembeddingsgo "github.com/mromero1591/openai-embeddings-go"
	"github.com/mromero1591/openai-embeddings-go/pinecone"
)

func FormatOpenAIEmbeddingToPineconeVector(inputs []string, keys []string, embeddings []openaiembeddingsgo.OpenAIEmbedding) {
	pineconeEmbeddings := make([]pinecone.PineconeVector, len(embeddings))
	for i, embedding := range embeddings {
		pineconeEmbeddings[i] = pinecone.PineconeVector{
			ID:     fmt.Sprintf("openAI-%s-%d", keys[i], embedding.Index),
			Values: embedding.Embedding,
			MetaData: pinecone.PineconeMetaData{
				Text: inputs[i],
			},
		}
	}
}

func EncodeTextToTokens(s string) int {
	// Split the string into tokens using whitespace as the delimiter
	tokens := strings.Fields(s)
	// Estimate the number of characters per token as 4
	charactersPerToken := 4
	// Calculate the total number of characters in the string
	totalCharacters := len(s)
	// Calculate the estimated number of tokens based on the given rule
	estimatedTokenCount := int(math.Ceil(float64(totalCharacters) / float64(charactersPerToken)))
	// Return the maximum of the estimated token count and the actual token count
	return int(math.Max(float64(estimatedTokenCount), float64(len(tokens))))
}
