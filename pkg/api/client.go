// Package api contains the HTTP client for the Jenkins instance
package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// JenkinsClient is a specialized version of the http.Client
type JenkinsClient struct {
	http.Client
	BaseURL             string
	authorizationHeader string
}

// NewJenkinsClient creates a new Jenkins Client
// This client implements a 15 second timeout
func NewJenkinsClient(baseURL, username, password string) *JenkinsClient {
	authorizationHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return &JenkinsClient{
		Client: http.Client{
			Timeout: 15 * time.Second,
		},
		BaseURL:             baseURL,
		authorizationHeader: fmt.Sprintf("Basic %s", authorizationHeader),
	}
}

// Do performs a HTTP request.
// This also adds the authorization header from values found in the environment variables
func (c *JenkinsClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", c.authorizationHeader)
	return c.Client.Do(req)
}

// Get is syntactic sugar for a HTTP Do.
// Needs to be reimplemented to get the benefits of having the headers
func (c *JenkinsClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
