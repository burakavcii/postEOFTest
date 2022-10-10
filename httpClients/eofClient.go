package httpClients

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type EOFHTTPClient struct {
	client http.Client
	url    *url.URL
}

func NewEOFHTTPClient() EOFHTTPClient {
	return EOFHTTPClient{client: http.Client{}}
}

func (c *EOFHTTPClient) EOFTest(URL string) error {
	parsedUrl, err := url.Parse(URL)
	if err != nil {
		log.Printf("error while parsing url: %s", err)
	}
	//eader: http.Header{"Connection": []string{"keep-alive"}}, Close: false
	httpRequest := http.Request{URL: parsedUrl, Method: "POST"}
	resp, err := c.client.Do(&httpRequest)
	if err != nil {
		//log.Printf("error while doing request: %s\n", err)
		return err
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		//log.Printf("error while reading body: %s\n", err)
		return err
	}
	return nil
}
