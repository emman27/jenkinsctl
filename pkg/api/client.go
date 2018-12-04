// Package api contains the HTTP client for the Jenkins instance
package api

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/pkg/errors"
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
	glog.Infof("Initializing Jenkins client with user: %s and host: %s", username, baseURL)
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
// In addition, the client also checks for non-2xx status codes and reports them as errors
func (c *JenkinsClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", c.authorizationHeader)
	glog.Infof("Calling %s", req.URL.String())
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "http request for %s not successful", req.URL.String())
	}
	err = checkStatusCode(resp)
	return resp, err
}

// Get is syntactic sugar for a HTTP Do.
// Needs to be reimplemented to get the benefits of having the headers
// Also automatically adds the base URL
func (c *JenkinsClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Post is syntactic sugar for a HTTP Do.
// Takes the benefits of the JenkinsClient.Get and replicates them here
func (c *JenkinsClient) Post(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.BaseURL+url, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func checkStatusCode(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		return fmt.Errorf("api call failed with status code %d", resp.StatusCode)
	}
	return nil
}
