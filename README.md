# go-harbor

Just for token based authentication.

Go library for Harbor v1.7.0 API

### Installation

Install using `go get github.com/XiaYinchang/harbor-go-sdk`.

### Usage

```go
// import package
import "github.com/XiaYinchang/harbor-go-sdk/harbor"

// create new client
authInfo := harbor.HarborAuth {
    APIURL:     "http://192.168.56.101/api",
	Token:  "GAAAJYFLFJKLSKHFSKLJFSLFJ",
}

client, err := harbor.NewClientWithToken(authInfo)
if err != nil {
    log.Fatal(err)
}

// get project by name
projectm, err := client.GetProjectByNmae("admin")
if err != nil {
    log.Fatal(err)
}

```
