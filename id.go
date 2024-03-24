package prettyld

import (
	"encoding/json"
	"errors"
	"reflect"
)

// ID represents a JSON-LD node.
//
// But the whole purprose of this type is to just be able to extract the ID from
// the JSON-LD node. So, it's just a string.
//
// This improves the readability of the code.
//
// Conveniently enough, when defining a struct for which it is pointing to an
// object, in your problem domain, sometimes, it is most likely guaranteed that
// you will only ever get one object. So, rather than defining the struct field
// as a slice, you can simply define it as the ID.
//
// For example:
//
//	type Person struct {
//		OtherNode ID `json:"https://example.com/ns#otherNode"`
//	}
//
// No need to define it as a slice.
//
// See: https://www.w3.org/TR/json-ld11-api/#idl-def-JsonLdNode
type ID string

var _ json.Marshaler = ID("")
var _ json.Unmarshaler = (*ID)(nil)

// MarshalJSON implements the json.Marshaler interface.
//
// But rather than marshal the string as a string, it marshals it as an object
// that has the key "@id" and the value as the string.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"@id": string(id),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It will either unmarshal a slice of objects that contain an `@id` key, or
// just an object that contains an `@id` key.
//
// Why not a string? The assumption is, we are intending to only parse expanded
// documents.
//
// THIS SHOULD NOT BE USED WITH UNEXPANDED DOCUMENTS!
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
