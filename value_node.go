package prettyld

import (
	"encoding/json"
	"errors"
	"reflect"
)

type ValueNode[V any] struct {
	Value    V      `json:"@value"`
	Language string `json:"@language,omitempty"`
	Type     string `json:"@type,omitempty"`
}

var _ json.Marshaler = ValueNode[int]{}
var _ json.Unmarshaler = (*ValueNode[int])(nil)

func (vn ValueNode[V]) MarshalJSON() ([]byte, error) {
	type t ValueNode[V]
	v := t(vn)
	return json.Marshal(v)
}

func (vn *ValueNode[V]) UnmarshalJSON(data []byte) error {
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
		return vn.UnmarshalJSON(b)
	}

	type t ValueNode[V]

	var vn2 t

	err = json.Unmarshal(data, &vn2)
	if err != nil {
		return err
	}

	*vn = ValueNode[V](vn2)

	return nil
}
