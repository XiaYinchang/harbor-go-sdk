package harbor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GetProjectByNmae(name string) (*ResProject, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:    "/projects",
		Method: http.MethodGet,
		Parameters: url.Values{
			"name": []string{
				name,
			},
		},
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resProjects []ResProject
	err = json.Unmarshal(resp.Body, &resProjects)

	if err != nil {
		return nil, err
	}
	if len(resProjects) <= 0 {
		return nil, fmt.Errorf("Error: no project named %s", name)
	}
	return &(resProjects[0]), nil
}
