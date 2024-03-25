package prettyld

import "github.com/piprate/json-gold/ld"

type WithContext struct {
	Context any
}

func (w WithContext) MarshalCompactJSONLD(source any, options *ld.JsonLdOptions) ([]byte, error) {
	return marshalCompactJSONLD(source, w, options)
}
