# AWS Signer

# Usage

Install the library as usual:

```bash 
go get github.com/jensskott/aws-signer
```

If you want to run the tests, just in case... :

```bash
cd $GOPATH/src/github.com/jensskott/aws-signer
go test -cover
```

For example, if you're using ElasticSearch with [@olivere](github.com/olivere)'s elastic library:

```go 
import (
    "github.com/jensskott/aws-signer"
    "github.com/aws/aws-sdk-go/service/elasticsearchservice"
    "gopkg.in/olivere/elastic.v3"
)

// Example For ElasticSearch

transport := signer.NewTransport(c, elasticsearchservice.ServiceName)

httpClient := &http.Client{
	Transport: transport,
}
// Use the client with Olivere's elastic client
client, err := elastic.NewClient(
    elastic.SetSniff(false),
    elastic.SetURL("your-aws-es-endpoint"),
    elastic.SetScheme("https"),
    elastic.SetHttpClient(httpClient),
)
```
