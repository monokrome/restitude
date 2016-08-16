package main

import (
	"errors"
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

func (res ExampleResource) GetItem(identifier string, r *http.Request) (interface{}, error) {
	return Example{Name: identifier}, nil
}

func (res ExampleResource) PostCollection(r *http.Request) (interface{}, error) {
	return nil, errors.New("POST is not a supported HTTP method.")
}

func main() {
	api := restitude.NewRestApi("/api/")
	api.RegisterResource(ExampleResource{})
	log.Print("Example service at http://", bindAddress, "/api/example/")
	log.Fatalln(http.ListenAndServe(bindAddress, nil))
}
