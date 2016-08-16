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
	DeleteItem(r *http.Request) (interface{}, error)
}

type getItemResource interface {
	GetItem(r *http.Request) (interface{}, error)
}

type headItemResource interface {
	HeadItem(r *http.Request) (interface{}, error)
}

type optionsItemResource interface {
	OptionsItem(r *http.Request) (interface{}, error)
}

type patchItemResource interface {
	PatchItem(r *http.Request) (interface{}, error)
}

type postItemResource interface {
	PostItem(r *http.Request) (interface{}, error)
}

type putItemResource interface {
	PutItem(r *http.Request) (interface{}, error)
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
