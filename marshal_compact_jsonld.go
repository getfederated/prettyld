package prettyld

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

var proc = ld.NewJsonLdProcessor()

func isMSA(source any) bool {
	t := reflect.TypeOf(source)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Map {
		if t.Key().Kind() != reflect.String {
			return false
		}
	} else {
		return false
	}
	return true
}

func isValidObject(source any) bool {
	// TODO: perhaps it should be a red flag if the source is a slice. Should we
	//   not also check the values inside sub structs?
	//
	// Or will the marshalCompactJSONLD do that recursively?

	t := reflect.ValueOf(source)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// If not struct, then it should be false…
	if t.Kind() != reflect.Struct {
		// Unless…

		if t.Kind() == reflect.Slice {
			for i := 0; i < t.Len(); i++ {
				if !isValidObject(t.Index(i).Interface()) {
					return false
				}
			}
		} else if !isMSA(source) {
			return false
		}
	}
	return true
}

type Context any

// marshalCompactJSONLD marshals the source object into a compact JSON-LD
// document, as a byte slice.
func marshalCompactJSONLD(source any, context Context, options *ld.JsonLdOptions) ([]byte, error) {
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

	var msa any
	err = json.Unmarshal(b, &msa)
	if err != nil {
		return nil, err
	}

	inter, err := proc.Compact(msa, context, options)
	if err != nil {
		return nil, err
	}

	return json.Marshal(inter)
}
