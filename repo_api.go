package harbor

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Client) GetRepositoriesByProjctNmae(name string) ([]Repository, error) {
	projectInfo, err := c.GetProjectByNmae(name)
	if err != nil {
		return nil, err
	}
	resp, err := c.DoRequest(KeyRequest{
		URL:    "/repositories",
		Method: http.MethodGet,
		Parameters: url.Values{
			"project_id": []string{
				string(projectInfo.ProjectId),
			},
		},
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resRepositories []Repository
	err = json.Unmarshal(resp.Body, &resRepositories)

	if err != nil {
		return nil, err
	}
	return resRepositories, nil
}
