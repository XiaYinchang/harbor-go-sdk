package harbor

import (
	"net/http"
	"net/url"
)

const X_SUBJECT_TOKEN_HEADER = "X-Subject-Token"
const X_AUTH_TOKEN = "X-Auth-Token"

type HarborAuth struct {
	AuthURL     string
	APIURL      string
	RegistryUrl string
	APIVersion  string
	DomainName  string
	ProjectName string
	UserName    string
	UserId      string
	Password    string
	Token       string
}

type SystemInfo struct {
	RegistryUrl string `json:"registry_url"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ResTokenBody struct {
	Token ResToken `json:"token"`
}

type ResToken struct {
	User ResUser `json:"user"`
}

type ResUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type KeyRequest struct {
	URL          string
	Method       string
	Headers      map[string]string
	Parameters   url.Values
	Body         []byte
	OkStatusCode int
}

type KeyResponse struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
}

type ResProject struct {
	ProjectId         string          `json:"project_id"`
	OwnerId           string          `json:"owner_id"`
	Name              string          `json:"name"`
	CreateTime        string          `json:"creation_time"`
	UpdateTime        string          `json:"update_time"`
	Deleted           bool            `json:"deleted"`
	OwnerName         string          `json:"owner_name"`
	Togglable         bool            `json:"togglabe"`
	CurrentUserRoleId int             `json:"current_user_role_id"`
	RepoCount         int             `json:"repo_count"`
	ChartCount        int             `json:"chart_count"`
	Metadata          ProjectMetadata `json:"metadata"`
}

type ProjectMetadata struct {
	Public             bool   `json:"public"`
	EnableContentTrust string `json:"enable_content_trust"`
	PreventVul         string `json:"prevent_vul"`
	Severity           string `json:"severity"`
	AutoScan           string `json:"auto_scan"`
}

type Repository struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ProjectId   int     `json:"project_id"`
	Description string  `json:"description"`
	PullCount   int     `json:"pull_count"`
	StarCount   int     `json:"star_count"`
	TagsCount   int     `json:"tags_count"`
	Labels      []Label `json:"labels"`
	CreateTime  string  `json:"creation_time"`
	UpdateTime  string  `json:"update_time"`
}

type Label struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Scope       string `json:"scope"`
	ProjectId   int    `json:"project_id"`
	CreateTime  string `json:"creation_time"`
	UpdateTime  string `json:"update_time"`
	Deleted     bool   `json:"deleted"`
}

type Tag struct {
	Name         string `json:"name"`
	CreateTime   string `json:"created"`
	Architecture string `json:"architecture"`
	Author       string `json:"author"`
	Os           string `json:"os"`
	Size         int    `json:"size"`
	Digest       string `json:"digest"`
}

type CreateProjectBody struct {
	ProjectName string          `json:"project_name"`
	Metadata    ProjectMetadata `json:"metadata"`
}
