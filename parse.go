package prettyld

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

// This is just a helper type to make it easier to work with expanded JSON-LD
// documents.
type LDNodesList []any

var _ json.Unmarshaler = (*LDNodesList)(nil)

func isString(b any) bool {
	valueOf := reflect.ValueOf(b)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	if valueOf.Kind() == reflect.String {
		return true
	}

	return false
}

func isByteSlice(b any) bool {
	valueOf := reflect.ValueOf(b)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	if valueOf.Kind() == reflect.Slice {
		if reflect.TypeOf(b).Elem().Kind() == reflect.Uint8 {
			return true
		}
	}

	return false
}

func getString(b any) (string, error) {
	valueOf := reflect.ValueOf(b)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	if valueOf.Kind() == reflect.String {
		return valueOf.String(), nil
	}

	return "", errors.New("expected a string")
}

func getByteSlice(b any) ([]byte, error) {
	valueOf := reflect.ValueOf(b)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	if valueOf.Kind() == reflect.Slice {
		if reflect.TypeOf(b).Elem().Kind() == reflect.Uint8 {
			return valueOf.Interface().([]byte), nil
		}
	}

	return nil, errors.New("expected a byte slice")
}

// Parse parses something in hopes that it is a JSON-LD document, and returns
// an LDNodesList, that will enable you to later unmarshal it into a struct of
// your choosing.
func Parse(b any, options *ld.JsonLdOptions) (LDNodesList, error) {
	if options == nil {
		options = ld.NewJsonLdOptions("")
	}

	isStr := isString(b)
	isB := isByteSlice(b)

	var bytes []byte

	if isStr {
		str, err := getString(b)
		if err != nil {
			return nil, err
		}
		bytes = []byte(str)
	} else if isB {
		str, err := getByteSlice(b)
		if err != nil {
			return nil, err
		}
		bytes = str
	}

	proc := ld.NewJsonLdProcessor()

	if isStr || isB {
		var a any

		err := json.Unmarshal(bytes, &a)
		if err != nil {
			return nil, err
		}

		i, err := proc.Expand(a, options)
		if err != nil {
			return nil, err
		}

		return LDNodesList(i), nil
	}

	expanded, err := proc.Expand(b, options)
	if err != nil {
		return nil, err
	}

	bytes, err = json.Marshal(expanded)
	if err != nil {
		return nil, err
	}

	return Parse(bytes, options)
}

func (p LDNodesList) UnmarshalTo(dest any) error {
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Slice {
		if len(p) != 1 {
			return errors.New("expected exactly 1 node in the list of parsed nodes")
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

// UnmarshalJSON is used for unmarshal an expanded document into something that
// is easier to work with.
//
// Should NOT be used with an unexpanded JSON object!
func (p *LDNodesList) UnmarshalJSON(b []byte) error {
	var a []any

	fmt.Println(string(b))

	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}

	*p = LDNodesList(a)

	return nil
}

// Iterate returns a channel that you can range over to get the nodes in the
// list.
func (p LDNodesList) Iterate() <-chan UnknownNode {
	ch := make(chan UnknownNode)

	go func() {
		defer close(ch)
		for _, v := range p {
			u, ok := v.(UnknownNode)
			if ok {
				ch <- UnknownNode(u)
			}
		}
	}()

	return ch
}
