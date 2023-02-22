# OPENAI-EMBEDDINGS-GO

## OpenAI and Pinecone API Clients

This repository contains two Client for calling OpenAI and Pinecone API endpoints respectively. These clients simplify the process of calling these endpoints and can be easily integrated into your existing codebase.

### OpenAI Client

The OpenAI wrapper provides an easy-to-use interface for calling OpenAI embedding endpoints. With this client, you can quickly generate embeddings for text data using OpenAI's powerful language models.

### Usage

To use the OpenAI client, you need to create an instance of the `Client` struct and set your OpenAI API key and embedding model name. Here's an example:

```go
import openaiembeddingsgo "github.com/mromero1591/openai-embeddings-go"

//create client and set model
openAIClient := openaiembeddingsgo.NewClient(apiKey)
openAIClient.WithModel(ADA002)

//create embeddings
response, err := CreateOpenAIEmbeddings("The fox ran over the hill.")
if err != nil {
    // handle error
}

```

The CreateOpenAIEmbeddings method of the OpenAIClient struct takes a string input and returns an OpenAIEmbeddingResponse struct containing the generated embeddings.

## Pinecone Client

The Pinecone Client provides an easy-to-use interface for calling Pinecone API endpoints. With this client, you can easily index and search high-dimensional vector data using Pinecone's vector database.

## Usage

To use the Pinecone client, you need to create an instance of the PineconeClient struct and set your Pinecone API key and Pinecone Environment and ProjectId. Here's an example:

```go
import "github.com/mromero1591/openai-embeddings-go/pinecone"

pineconeClient := pinecone.NewClient(env, apiKey, projectID)

//create an index
indexConfig := pinecone.PineconeIndexConfig{
    Name: "test-index",
    Dimension: 1024
}

if err := pineconeClient.CreatePineconeIndex(indexConfig) {
    //handle err
}

```

The Pinecone Client has 3 methods to work with Indexes. `ListPineconeIndex` returns a list of your Pinecone indexes, `UpsertPineconeIndex` writes vectors into a namespace. If a new value is upserted for an existing vector id, it will overwrite the previous value. `QueryPineconeIndex` searches a namespace, using a query vector. It retrieves the ids of the most similar items in a namespace, along with their similarity scores.

## Contributing

If you would like to contribute to this repository, please fork the repository, make your changes, and submit a pull request. We welcome contributions of all types, including bug fixes, feature additions, and documentation improvements.

## License

This repository is licensed under the MIT License. See the [LICENSE](https://chat.openai.com/LICENSE) file for more information.
