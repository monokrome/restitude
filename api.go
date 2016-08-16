package restitude

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

const MatchingResourceNotFound = "No resource found matching the given request."

var requestMethods = []string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"}

type restApiItemHandlerStore map[string]func(identifier string, r *http.Request) (interface{}, error)
type restApiCollectionHandlerStore map[string]func(r *http.Request) (interface{}, error)

type restApi struct {
	Prefix      string
	Serializers map[string]restSerializer

	itemResources       map[string]restApiItemHandlerStore
	collectionResources map[string]restApiCollectionHandlerStore
}

type ErrorResponse struct {
	Message string
}

// Get the API base name for the given resource
func getBaseName(iface interface{}) string {
	if resource, ok := iface.(resource); ok {
		return resource.BaseName()
	}

	// If no BaseName is provided, guess based on the type's name
	typeName := reflect.TypeOf(iface).Name()
	if len(typeName) > 8 && typeName[len(typeName)-8:] == "Resource" {
		typeName = typeName[:len(typeName)-8]
	}
	return strings.ToLower(typeName)
}

func NewRestApi(prefix string) *restApi {
	log.Print("Creating new REST API at ", prefix)

	itemResources := make(map[string]restApiItemHandlerStore)
	collectionResources := make(map[string]restApiCollectionHandlerStore)

	for _, method := range requestMethods {
		itemResources[method] = make(restApiItemHandlerStore)
		collectionResources[method] = make(restApiCollectionHandlerStore)
	}

	api := &restApi{
		Prefix:              prefix,
		Serializers:         getDefaultSerializers(),
		itemResources:       itemResources,
		collectionResources: collectionResources,
	}

	http.HandleFunc(api.Prefix, api.onRequestReceived)

	return api
}

// Allows registering of resources to specific APIs
func (api *restApi) RegisterResource(iface interface{}) {
	baseName := getBaseName(iface)

	if handler, ok := iface.(deleteItemResource); ok {
		api.itemResources["DELETE"][baseName] = handler.DeleteItem
	}

	if handler, ok := iface.(getItemResource); ok {
		api.itemResources["GET"][baseName] = handler.GetItem
	}

	if handler, ok := iface.(headItemResource); ok {
		api.itemResources["HEAD"][baseName] = handler.HeadItem
	}

	if handler, ok := iface.(patchItemResource); ok {
		api.itemResources["PATCH"][baseName] = handler.PatchItem
	}

	if handler, ok := iface.(postItemResource); ok {
		api.itemResources["POST"][baseName] = handler.PostItem
	}

	if handler, ok := iface.(putItemResource); ok {
		api.itemResources["PUT"][baseName] = handler.PutItem
	}

	if handler, ok := iface.(deleteCollectionResource); ok {
		api.collectionResources["DELETE"][baseName] = handler.DeleteCollection
	}

	if handler, ok := iface.(getCollectionResource); ok {
		api.collectionResources["GET"][baseName] = handler.GetCollection
	}

	if handler, ok := iface.(headCollectionResource); ok {
		api.collectionResources["HEAD"][baseName] = handler.HeadCollection
	}

	if handler, ok := iface.(patchCollectionResource); ok {
		api.collectionResources["PATCH"][baseName] = handler.PatchCollection
	}

	if handler, ok := iface.(postCollectionResource); ok {
		api.collectionResources["POST"][baseName] = handler.PostCollection
	}

	if handler, ok := iface.(putCollectionResource); ok {
		api.collectionResources["PUT"][baseName] = handler.PutCollection
	}
}

func (api *restApi) handleItem(baseName string, identifier string, r *http.Request) (interface{}, error) {
	if resources, ok := api.itemResources[r.Method]; ok {
		if handler, ok := resources[baseName]; ok {
			response, err := handler(identifier, r)
			return response, err
		}
	}

	return nil, errors.New(MatchingResourceNotFound)
}

func (api *restApi) handleCollection(baseName string, r *http.Request) (interface{}, error) {
	if resources, ok := api.collectionResources[r.Method]; ok {
		if handler, ok := resources[baseName]; ok {
			response, err := handler(r)
			return response, err
		}
	}

	return nil, errors.New(MatchingResourceNotFound)
}

func (api *restApi) getResponseSerializer(r *http.Request) restSerializer {
	accept := r.Header.Get("Accept")

	// TODO: Don't use split. It's excessive here.
	for _, contentType := range strings.Split(accept, ",") {
		// TODO: Support proper ordering with q=
		contentType = strings.Split(contentType, ";")[0]

		if serializer, ok := api.Serializers[contentType]; ok {
			return serializer
		}
	}

	return json.Marshal
}

// Handle routing of requests to their resources
func (api *restApi) onRequestReceived(w http.ResponseWriter, r *http.Request) {
	log.Print("Received ", r.Method, " request at: ", r.RequestURI)

	trunctedString := strings.TrimRight(r.RequestURI[len(api.Prefix):], "/")

	// NOTE: Split/Join is excessive here and doesn't help performance.
	parts := strings.Split(trunctedString, "/")
	identifier := strings.Join(parts[:], "/")

	serialize := api.getResponseSerializer(r)

	var response interface{}
	var err error

	if len(parts) > 1 {
		response, err = api.handleItem(parts[0], identifier, r)
	} else if parts[0] == "" {
		err = errors.New("Support for generating schemas is not yet implemented.")
	} else {
		response, err = api.handleCollection(parts[0], r)
	}

	if err != nil {
		response = ErrorResponse{
			Message: fmt.Sprintf("%s", err),
		}
	}

	output, err := serialize(response)

	if err != nil {
		// TODO: Handle this error case
		log.Print(err)
		return
	}

	w.Write(output)
}
