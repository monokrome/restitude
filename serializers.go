package restitude

import (
	"encoding/json"
	"encoding/xml"
)

type restSerializer func(v interface{}) ([]byte, error)

func serializeToJSON(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func serializeToXML(v interface{}) ([]byte, error) {
	data, err := xml.Marshal(v)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func getDefaultSerializers() map[string]restSerializer {
	serializers := make(map[string]restSerializer)
	serializers["application/json"] = serializeToJSON
	serializers["application/xml"] = serializeToXML
	return serializers
}
