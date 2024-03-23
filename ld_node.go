package prettyld

import (
	"encoding/json"
	"errors"
	"reflect"
)

type LDNode struct {
	node any
}

var _ json.Marshaler = LDNode{}
var _ json.Unmarshaler = (*LDNode)(nil)

func (n LDNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.node)
}

func (n *LDNode) UnmarshalJSON(data []byte) error {
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
		return n.UnmarshalJSON(b)
	}

	var n2 any

	err = json.Unmarshal(data, &n2)
	if err != nil {
		return err
	}

	*n = LDNode{node: n2}

	return nil
}
