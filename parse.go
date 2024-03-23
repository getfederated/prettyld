package prettyld

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

type Parsed struct {
	proc    *ld.JsonLdProcessor
	options *ld.JsonLdOptions
	payload []any
}

func Parse(b []byte, options *ld.JsonLdOptions) (*Parsed, error) {
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

	return &Parsed{proc: proc, payload: i}, nil
}

func (p *Parsed) Unmarshal(dest any) error {
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Slice {
		if len(p.payload) != 1 {
			return errors.New("expected a maximum of 1 node in the list of parsed nodes")
		}
		b, err := json.Marshal(p.payload[0])
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &dest)
		if err != nil {
			return err
		}
		return nil
	}

	b, err := json.Marshal(p.payload)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &dest)
}
