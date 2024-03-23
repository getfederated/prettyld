package prettyld

import (
	"encoding/json"
)

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
