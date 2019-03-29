package harbor

import (
	"encoding/json"
	"net/http"
)

func (c *Client) CreateUser(name string) error {
	body, err := json.Marshal(UserInfo{
		Username: name,
		Password: "test",
		Email:    name + "@example.com",
	})
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/users",
		Method:       http.MethodPost,
		Body:         body,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return err
	}
	return nil
}
