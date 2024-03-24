package prettyld

import "encoding/json"

type UnknownNode map[string]any

// GetObject returns the object associated with the given predicate.
func (u UnknownNode) GetObject(predicate string) []UnknownNode {
	if predicate[0] == '@' {
		return []UnknownNode{}
	}

	object, ok := u[predicate]
	if !ok {
		return []UnknownNode{}
	}
	node, ok := object.(map[string]any)
	if !ok {
		return []UnknownNode{}
	}
	return []UnknownNode{node}
}

// UnmarshalTo unmarshals the node into the given destination. It is analogous
// to Parser.UnmarshalTo.
func (u UnknownNode) UnmarshalTo(dest any) error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &dest)
}
