# restitude
A tiny library for creating RESTful services in Go.


## Usage

Here's an example of creating a simple API with a manually provided serializer
for msgpack support:

```go
package main

import (
	"errors"
	"github.com/menghan/msgpack"
	"github.com/monokrome/restitude"
	"log"
	"net/http"
)

const bindAddress = "127.0.0.1:9876"

type Example struct {
	Name string
}

type ExampleResource struct{}

func (res ExampleResource) GetCollection(r *http.Request) (interface{}, error) {
	return []Example{Example{Name: "Bailey"}}, nil
}

func (res ExampleResource) GetItem(r *http.Request) (interface{}, error) {
	return Example{Name: "Bailey"}, nil
}

func (res ExampleResource) PostCollection(r *http.Request) (interface{}, error) {
	return nil, errors.New("POST is not a supported HTTP method.")
}

func main() {
	api := restitude.NewRestApi("/api/")
    api.Serializers["application/msgpack"] = msgpack.Marshal
	api.RegisterResource(ExampleResource{})
	log.Print("Example service at http://", bindAddress, "/api/example/")
	log.Fatalln(http.ListenAndServe(bindAddress, nil))
}
```
