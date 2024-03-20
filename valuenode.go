package prettyld

type ValueNode[V any] struct {
	Value    V      `json:"@value"`
	Language string `json:"@language,omitempty"`
}
