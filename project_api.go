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

func (c *Client) CheckProjectExist(name string) (bool, error) {
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
		return true, err
	}
	var resProjects []ResProject
	err = json.Unmarshal(resp.Body, &resProjects)

	if err != nil {
		return true, err
	}
	if len(resProjects) <= 0 {
		return false, nil
	}
	return true, nil
}

func (c *Client) CreateProject(projectName, ownerName string) error {
	body, err := json.Marshal(CreateProjectBody{
		ProjectName: projectName,
		OwnerName:   ownerName,
		Metadata: map[string]string{
			"public": "false",
		},
	})
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/projects",
		Method:       http.MethodPost,
		Body:         body,
		OkStatusCode: http.StatusCreated,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteProject(name string) error {
	retRepos, err := c.GetRepositoriesByProjectName(name)
	if err != nil {
		return err
	}
	for _, v := range retRepos {
		err = c.DeleteRepo(v.Name)
		if err != nil {
			return err
		}
	}
	projectInfo, err := c.GetProjectByNmae(name)
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/projects/" + fmt.Sprintf("%v", projectInfo.ProjectId),
		Method:       http.MethodDelete,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetUserProjects(userName string) ([]ResProject, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:    "/projects",
		Method: http.MethodGet,
		Parameters: url.Values{
			"owner": []string{
				userName,
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
	return resProjects, nil
}
