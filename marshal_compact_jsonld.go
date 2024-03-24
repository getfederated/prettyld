package prettyld

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

var proc = ld.NewJsonLdProcessor()

func isValidObject(source any) bool {
	// TODO: perhaps it should be a red flag if the source is a slice. Should we
	//   not also check the values inside sub structs?
	//
	// Or will the MarshalCompactJSONLD do that recursively?

	t := reflect.TypeOf(source)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		if t.Kind() == reflect.Slice {
			v := reflect.ValueOf(source)
			for i := 0; i < v.Len(); i++ {
				if v.Index(i).Kind() != reflect.Struct {
					if v.Index(i).Kind() != reflect.Interface {
						return false
					}
					value := v.Index(i).Interface()
					if reflect.TypeOf(value).Kind() != reflect.Struct {
						return false
					}
				}
			}
		} else {
			return false
		}
	}
	return true
}

// MarshalCompactJSONLD marshals the source object into a compact JSON-LD
// document, as a byte slice.
func MarshalCompactJSONLD(source any, options *ld.JsonLdOptions) ([]byte, error) {
	if options == nil {
		options = ld.NewJsonLdOptions("")
	}

	if !isValidObject(source) {
		return nil, fmt.Errorf("source must be a struct, pointer to a struct, or a slice of structs")
	}

	b, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(source)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Slice {
		var msa []any
		err = json.Unmarshal(b, &msa)
		if err != nil {
			return nil, err
		}

		inter, err := proc.Compact(msa, nil, options)
		if err != nil {
			return nil, err
		}

		return json.Marshal(inter)
	}

	var msa any
	err = json.Unmarshal(b, &msa)
	if err != nil {
		return nil, err
	}

	inter, err := proc.Compact(msa, nil, options)
	if err != nil {
		return nil, err
	}

	return json.Marshal(inter)
}
