package harbor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AuthInfo *HarborAuth
}

func NewClientWithToken(authInfo *HarborAuth) (*Client, error) {
	if authInfo.Token == "" {
		return nil, fmt.Errorf("missing token")
	}
	client := Client{AuthInfo: authInfo}
	systemInfo, err := client.GetRegistry()
	if err != nil {
		return nil, err
	}
	client.AuthInfo.RegistryUrl = systemInfo.RegistryUrl
	return &client, nil
}

func NewClientWithPassword(authInfo *HarborAuth) (*Client, error) {
	if authInfo.UserName == "" || authInfo.Password == "" {
		return nil, fmt.Errorf("missing username or password")
	}
	client := Client{AuthInfo: authInfo}
	systemInfo, err := client.GetRegistry()
	if err != nil {
		return nil, err
	}
	client.AuthInfo.RegistryUrl = systemInfo.RegistryUrl
	return &client, nil
}

func (c *Client) DoRequest(r KeyRequest) (KeyResponse, error) {
	client := &http.Client{}
	reqUrl := c.AuthInfo.APIURL + r.URL + "?" + r.Parameters.Encode()
	req, err := http.NewRequest(r.Method, reqUrl, bytes.NewBuffer(r.Body))
	if err != nil {
		return KeyResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.AuthInfo.Token != "" {
		req.Header.Set(X_AUTH_TOKEN, c.AuthInfo.Token)
	}
	if c.AuthInfo.UserName != "" && c.AuthInfo.Password != "" {
		req.SetBasicAuth(c.AuthInfo.UserName, c.AuthInfo.Password)
	}
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

func (c *Client) GetRegistry() (*SystemInfo, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/systeminfo",
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var systemInfo SystemInfo
	err = json.Unmarshal(resp.Body, &systemInfo)

	if err != nil {
		return nil, err
	}
	return &systemInfo, nil
}
