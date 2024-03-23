package prettyld

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

type LDNodes []any

func Parse(b []byte, options *ld.JsonLdOptions) (LDNodes, error) {
	if options == nil {
		options = ld.NewJsonLdOptions("")
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

	return LDNodes(i), nil
}

func (p LDNodes) UnmarshalTo(dest any) error {
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

func (p LDNodes) Items() <-chan *LDNodes {
	c := make(chan *LDNodes)
	go func() {
		defer close(c)
		for _, item := range p {
			c <- &LDNodes{item}
		}
	}()
	return c
}

// UnmarshalJSON is used unmarshal an expanded document into something that is
// easier to work with.
//
// Should NOT be used with an unexpanded JSON object!
func (p *LDNodes) UnmarshalJSON(b []byte) error {
	var a []any

	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}

	*p = LDNodes(a)

	return nil
}
