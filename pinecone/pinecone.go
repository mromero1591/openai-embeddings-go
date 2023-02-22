package pinecone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	Environment string `json:"environment"`
	APIKey      string `json:"apiKey"`
	ProjectID   string `json:"projectId"`
}

func NewClient(env string, apiKey string, projectID string) Client {
	return Client{
		Environment: env,
		APIKey:      apiKey,
		ProjectID:   projectID,
	}
}

func (p *PineconeIndexConfig) SetPineconeIndexDefaults() {
	if p.Metric == "" {
		p.Metric = "cosine"
	}
	if p.Pods == 0 {
		p.Pods = 1
	}
	if p.Replicas == 0 {
		p.Replicas = 1
	}
	if p.Shards == 0 {
		p.Shards = 1
	}
	if p.PodType == "" {
		p.PodType = "p1.x1"
	}
}

func (c Client) CreatePineconeIndex(indexConfig PineconeIndexConfig) error {
	if indexConfig.Name == "" {
		return errors.New("missing Name in index config")
	}

	if indexConfig.Dimension == 0 {
		return errors.New("missing dimension in index config")
	}

	//setup defaults if missing
	indexConfig.SetPineconeIndexDefaults()

	url := fmt.Sprintf("https://controller.%s.pinecone.io/databases", c.Environment)

	requestBodyBytes, _ := json.Marshal(indexConfig)
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	request.Header.Set("accept", "text/plain")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Api-Key", c.APIKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func (c Client) ListPineconeIndexes() (PineconeListIndexResponse, error) {
	url := fmt.Sprintf("https://controller.%s.pinecone.io/databases", c.Environment)

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Api-Key", c.APIKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return PineconeListIndexResponse{}, err
	}
	defer response.Body.Close()

	var pineconeListIndexResponse PineconeListIndexResponse
	if err := json.NewDecoder(response.Body).Decode(&pineconeListIndexResponse); err != nil {
		return PineconeListIndexResponse{}, err
	}

	return pineconeListIndexResponse, nil
}

func (c Client) UpsertPineconeIndex(index string, requestBody PineconeUpsertIndexRequest) (PineconeUpsertIndexResponse, error) {
	url := fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io/vectors/upsert", index, c.ProjectID, c.Environment)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Api-Key", c.APIKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return PineconeUpsertIndexResponse{}, err
	}
	defer response.Body.Close()

	var pineconeUpsertIndexResponse PineconeUpsertIndexResponse
	if err := json.NewDecoder(response.Body).Decode(&pineconeUpsertIndexResponse); err != nil {
		return PineconeUpsertIndexResponse{}, err
	}

	return pineconeUpsertIndexResponse, nil
}

func (c Client) QueryPineconeIndex(index string, requestBody PineconeQueryIndexRequest) (PineconeQueryIndexResponse, error) {
	url := fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io/vectors/upsert", index, c.ProjectID, c.Environment)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Api-Key", c.APIKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return PineconeQueryIndexResponse{}, err
	}
	defer response.Body.Close()

	var pineconeQueryIndexResponse PineconeQueryIndexResponse
	if err := json.NewDecoder(response.Body).Decode(&pineconeQueryIndexResponse); err != nil {
		return PineconeQueryIndexResponse{}, err
	}

	return pineconeQueryIndexResponse, nil
}
