package restitude

import (
	"net/http"
)

// Interface for getting the base path to a RESTful resource
type resource interface {
	BaseName() string
}

// Interface for resources which are able to work with individual entities
type deleteItemResource interface {
	DeleteItem(identifier string, r *http.Request) (interface{}, error)
}

type getItemResource interface {
	GetItem(identifier string, r *http.Request) (interface{}, error)
}

type headItemResource interface {
	HeadItem(identifier string, r *http.Request) (interface{}, error)
}

type optionsItemResource interface {
	OptionsItem(identifier string, r *http.Request) (interface{}, error)
}

type patchItemResource interface {
	PatchItem(identifier string, r *http.Request) (interface{}, error)
}

type postItemResource interface {
	PostItem(identifier string, r *http.Request) (interface{}, error)
}

type putItemResource interface {
	PutItem(identifier string, r *http.Request) (interface{}, error)
}

// Interface for resources which are able to work with collections of entities
type deleteCollectionResource interface {
	DeleteCollection(r *http.Request) (interface{}, error)
}

type getCollectionResource interface {
	GetCollection(r *http.Request) (interface{}, error)
}

type headCollectionResource interface {
	HeadCollection(r *http.Request) (interface{}, error)
}

type optionsCollectionResource interface {
	OptionsCollection(r *http.Request) (interface{}, error)
}

type patchCollectionResource interface {
	PatchCollection(r *http.Request) (interface{}, error)
}

type postCollectionResource interface {
	PostCollection(r *http.Request) (interface{}, error)
}

type putCollectionResource interface {
	PutCollection(r *http.Request) (interface{}, error)
}
