package prettyld

import (
	"encoding/json"
)

// String is a type that represents a string value in a JSON-LD document.
//
// In reality, string is a convenience type that is intended for translating
// a `ValueNode` into a string. This is useful when you know that a value is
// a string and you want to avoid the overhead of working with a `ValueNode`.
type String string

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s String) MarshalJSON() ([]byte, error) {
	value := ValueNode[string]{
		Value: string(s),
	}
	return json.Marshal(value)
}

func (s *String) UnmarshalJSON(data []byte) error {
	v := ValueNode[string]{Value: string(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = String(v.Value)
	return nil
}
