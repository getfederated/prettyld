package prettyld

import "github.com/piprate/json-gold/ld"

type ContextInstance struct {
	context any
}

func WithContext(context any) ContextInstance {
	return ContextInstance{context}
}

func (w ContextInstance) MarshalCompactJSONLD(source any, options *ld.JsonLdOptions) ([]byte, error) {
	return marshalCompactJSONLD(source, w.context, options)
}
