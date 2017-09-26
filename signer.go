package signer

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

// AWSSigningTransport implements http.RoundTripper. When set as Transport of http.Client, it signs HTTP requests with the latest AWS v4 signer and logs the requests
// No field is mandatory, but you can provide your own Transport or contextLogger by setting the Transport or Logger property.
type AWSSigningTransport struct {
	Transport      http.RoundTripper
	awsServiceName string
	awsSigner      *v4.Signer
	awsRegion      string
}

// NewTransport creates a new instance of the AWSSigningTransport
func NewTransport(c *credentials.Credentials, region, serviceName string) *AWSSigningTransport {
	return &AWSSigningTransport{
		awsSigner:      v4.NewSigner(c),
		awsRegion:      region,
		awsServiceName: serviceName,
	}
}

func readAndReplaceBody(req *http.Request) []byte {
	if req.Body == nil {
		return []byte{}
	}
	payload, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewReader(payload))
	return payload
}

// RoundTrip is the core part of this module and implements http.RoundTripper.
// Executes HTTP request with request/response logging.
// Signs the requests with AWS's v4 signer
func (t *AWSSigningTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	payload := bytes.NewReader(readAndReplaceBody(req))
	_, err := t.awsSigner.Sign(req, payload, t.awsServiceName, t.awsRegion, time.Now())
	if err != nil {
		return nil, err
	}

	resp, err := t.transport().RoundTrip(req)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func (t *AWSSigningTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
