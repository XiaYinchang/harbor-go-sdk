package harbor

import (
	"net/http"
	"net/url"
	"time"
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
	UserId       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Realname     string `json:"realname"`
	Rolename     string `json:"role_name"`
	Role         int    `json:"role_id"`
	HasAdminRole bool   `json:"has_admin_role"`
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
	ProjectId         int               `json:"project_id"`
	OwnerId           int               `json:"owner_id"`
	Name              string            `json:"name"`
	CreateTime        time.Time         `json:"creation_time"`
	UpdateTime        time.Time         `json:"update_time"`
	Deleted           bool              `json:"deleted"`
	OwnerName         string            `json:"owner_name"`
	Togglable         bool              `json:"togglabe"`
	CurrentUserRoleId int               `json:"current_user_role_id"`
	RepoCount         int64             `json:"repo_count"`
	ChartCount        uint64            `json:"chart_count"`
	Metadata          map[string]string `json:"metadata"`
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
	Id          int32  `json:"id"`
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
	ProjectName string            `json:"project_name"`
	OwnerName   string            `json:"owner_name"`
	Metadata    map[string]string `json:"metadata"`
}
