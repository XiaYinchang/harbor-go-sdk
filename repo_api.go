package harbor

import (
	"encoding/json"
	"fmt"
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
				fmt.Sprintf("%v", projectInfo.ProjectId),
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

func (c *Client) DeleteTag(repoNameWithProjectName string, tagName string) error {
	_, err := c.DoRequest(KeyRequest{
		URL:          "/repositories/" + repoNameWithProjectName + "/tags/" + tagName,
		Method:       http.MethodDelete,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteRepo(repoNameWithProjectName string) error {
	_, err := c.DoRequest(KeyRequest{
		URL:          "/repositories/" + repoNameWithProjectName,
		Method:       http.MethodDelete,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetLabelsByRepositoryName(repoName string) ([]Label, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/repositories/" + repoName + "/labels",
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resLabels []Label
	err = json.Unmarshal(resp.Body, &resLabels)

	if err != nil {
		return nil, err
	}
	return resLabels, nil
}

func (c *Client) RetagRepo(targetRepoName, srcImage, targetTag string) error {

	reqBody, err := json.Marshal(
		struct {
			Tag      string `json:"tag"`
			SrcImage string `json:"src_image"`
			Override bool   `json:"override"`
		}{
			targetTag,
			srcImage,
			true,
		})
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/repositories/" + targetRepoName + "/tags",
		Method:       http.MethodPost,
		Body:         reqBody,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return err
	}
	return nil
}
