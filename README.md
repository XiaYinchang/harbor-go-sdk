# go-harbor :key:

Go library for Harbor v1.7.0 API

### Installation

Install using `go get github.com/XiaYinchang/harbor-go-sdk`.

### Usage

```go
// import package
import "github.com/XiaYinchang/harbor-go-sdk/harbor"

// create new client
authInfo := harbor.HarborAuth {
    AuthURL:     "http://192.168.56.101:5000",
	APIVersion:  "v3",
	DomainName:  "Default",
	ProjectName: "admin",
	UserName:    "admin",
	Password:    "test",
}

client, err := harbor.NewClient(authInfo)
if err != nil {
    log.Fatal(err)
}

// get token
token := client.AuthInfo.Token

// get projects of user with userid
userProjects := client.UserProjects(client.AuthInfo.UserId)
```
