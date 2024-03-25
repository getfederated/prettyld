package prettyld

import "github.com/piprate/json-gold/ld"

type withContext struct {
	context any
}

func WithContext(context any) withContext {
	return withContext{context}
}

func (w withContext) MarshalCompactJSONLD(source any, options *ld.JsonLdOptions) ([]byte, error) {
	return marshalCompactJSONLD(source, w.context, options)
}
