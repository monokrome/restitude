package restitude

import (
	"encoding/json"
	"encoding/xml"
)

type restSerializer func(v interface{}) ([]byte, error)

func getDefaultSerializers() map[string]restSerializer {
	serializers := make(map[string]restSerializer)
	serializers["application/json"] = json.Marshal
	serializers["application/xml"] = xml.Marshal
	return serializers
}
