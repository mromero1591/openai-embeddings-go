package tools

import (
	"fmt"

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
