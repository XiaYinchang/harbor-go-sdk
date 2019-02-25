package harbor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AuthInfo *HarborAuth
}

func NewClientWithToken(authInfo HarborAuth) (*Client, error) {
	if authInfo.Token == "" {
		return nil, fmt.Errorf("missing token")
	}
	client := Client{AuthInfo: &authInfo}
	return &client, nil
}

func (c *Client) DoRequest(r KeyRequest) (KeyResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, c.AuthInfo.AuthURL+r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return KeyResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(X_AUTH_TOKEN, c.AuthInfo.Token)
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return KeyResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return KeyResponse{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		return KeyResponse{}, fmt.Errorf("Error: %s details: %s\n", resp.Status, body)
	}

	return KeyResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header}, nil
}

func (c *Client) doRequest(r KeyRequest) (KeyResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return KeyResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return KeyResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return KeyResponse{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		return KeyResponse{}, fmt.Errorf("Error: %s details: %s\n", resp.Status, body)
	}

	return KeyResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header}, nil
}
