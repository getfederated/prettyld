package prettyld

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

type LDNodesList []any

var _ json.Unmarshaler = (*LDNodesList)(nil)

func Parse(b any, options *ld.JsonLdOptions) (LDNodesList, error) {
	if options == nil {
		options = ld.NewJsonLdOptions("")
	}

	valueOf := reflect.ValueOf(b)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	if valueOf.Kind() == reflect.Slice {
		if reflect.TypeOf(b).Elem().Kind() == reflect.Uint8 {
			b, ok := valueOf.Interface().([]byte)
			if !ok {
				panic("expected a byte slice. This is a severe internal error")
			}

			var a any

			proc := ld.NewJsonLdProcessor()

			err := json.Unmarshal(b, &a)
			if err != nil {
				return nil, err
			}

			i, err := proc.Expand(a, options)
			if err != nil {
				return nil, err
			}

			return LDNodesList(i), nil
		}
	}

	byres, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	return Parse(byres, options)
}

func (p LDNodesList) UnmarshalTo(dest any) error {
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Slice {
		if len(p) != 1 {
			return errors.New("expected a maximum of 1 node in the list of parsed nodes")
		}
		b, err := json.Marshal(p[0])
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &dest)
		if err != nil {
			return err
		}
		return nil
	}

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &dest)
}

// UnmarshalJSON is used unmarshal an expanded document into something that is
// easier to work with.
//
// Should NOT be used with an unexpanded JSON object!
func (p *LDNodesList) UnmarshalJSON(b []byte) error {
	var a []any

	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}

	*p = LDNodesList(a)

	return nil
}
