package harbor

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Client) GetRepositoriesByProjectName(name string) ([]Repository, error) {
	projectInfo, err := c.GetProjectByNmae(name)
	if err != nil {
		return nil, err
	}
	resp, err := c.DoRequest(KeyRequest{
		URL:    "/repositories",
		Method: http.MethodGet,
		Parameters: url.Values{
			"project_id": []string{
				projectInfo.ProjectId,
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

func (c *Client) GetTagsByRepositoryName(nameWithProjectName string) ([]Tag, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:    "/repositories/" + nameWithProjectName + "/tags",
		Method: http.MethodGet,
		Parameters: url.Values{
			"detail": []string{
				"1",
			},
		},
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resTags []Tag
	err = json.Unmarshal(resp.Body, &resTags)

	if err != nil {
		return nil, err
	}
	return resTags, nil
}
