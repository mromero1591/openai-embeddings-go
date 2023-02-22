package pinecone

type (
	PineconeIndexConfig struct {
		Metric    string `json:"metric"`
		Pods      int    `json:"pods"`
		Replicas  int    `json:"replicas"`
		Shards    int    `json:"shards"`
		PodType   string `json:"pod_type"`
		Name      string `json:"name"`
		Dimension int    `json:"dimension"`
	}

	PineconeListIndexResponse struct {
		Indexes []string `json:"indexes"`
	}

	PineconeUpsertIndexRequest struct {
		Vectors   []PineconeVector `json:"vectors"`
		Namespace string           `json:"namespace"`
	}

	PineconeVector struct {
		Values   []float64        `json:"values"`
		ID       string           `json:"id"`
		MetaData PineconeMetaData `json:"metadata"`
	}

	PineconeUpsertIndexResponse struct {
		UpsertedCount int `json:"upsertedCount"`
	}

	PineconeQueryIndexRequest struct {
		Namespace       string    `json:"namespace"`
		TopK            int       `json:"topK"`
		IncludeMetadata bool      `json:"includeMetadata"`
		IncludeValues   bool      `json:"includeValues"`
		Vector          []float64 `json:"vector"`
	}

	PineconeQueryIndexResponse struct {
		Matches   []PineconeMatch `json:"matches"`
		Namespace string          `json:"namespace"`
	}

	PineconeMatch struct {
		ID       string           `json:"id"`
		Score    float64          `json:"score"`
		MetaData PineconeMetaData `json:"metadata"`
	}

	PineconeMetaData struct {
		Text string `json:"text"`
	}
)
