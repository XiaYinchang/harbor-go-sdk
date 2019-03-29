package harbor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) CreateUser(name string) error {
	body, err := json.Marshal(User{
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
		OkStatusCode: http.StatusCreated,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetUserByName(name string) (*User, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:    "/users",
		Method: http.MethodGet,
		Parameters: url.Values{
			"username": []string{
				name,
			},
		},
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resUsers []User
	err = json.Unmarshal(resp.Body, &resUsers)
	if err != nil {
		return nil, err
	}
	if len(resUsers) <= 0 {
		return nil, fmt.Errorf("Error: no user named %s", name)
	}
	return &(resUsers[0]), nil
}

func (c *Client) DeleteUser(name string) error {
	userProjects, err := c.GetUserProjects(name)
	if err != nil {
		return err
	}
	for _, project := range userProjects {
		err = c.DeleteProject(project.Name)
		if err != nil {
			return err
		}
	}
	userinfo, err := c.GetUserByName(name)
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/users/" + fmt.Sprintf("%v", userinfo.UserId),
		Method:       http.MethodDelete,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return err
	}
	return nil
}
