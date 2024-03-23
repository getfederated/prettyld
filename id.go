package prettyld

import (
	"encoding/json"
	"errors"
	"reflect"
)

type ID string

var _ json.Marshaler = ID("")
var _ json.Unmarshaler = (*ID)(nil)

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"@id": string(id),
	})
}

func (id *ID) UnmarshalJSON(data []byte) error {
	var a any

	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	value := reflect.ValueOf(a)

	if value.Kind() == reflect.Slice {
		if value.Len() != 1 {
			return errors.New("expected a maximum of 1 element in the slice")
		}
		b, err := json.Marshal(value.Index(0).Interface())
		if err != nil {
			return err
		}
		return id.UnmarshalJSON(b)
	}

	type idNode struct {
		ID string `json:"@id"`
	}

	var node idNode
	err = json.Unmarshal(data, &node)
	if err != nil {
		return err
	}

	*id = ID(node.ID)
	return nil
}
