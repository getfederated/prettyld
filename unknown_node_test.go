package prettyld

import "testing"

func TestGetObject(t *testing.T) {
	t.Run("get object", func(t *testing.T) {
		s := UnknownNode{
			"https://example.com/ns#name": map[string]any{
				"@value": "John Doe",
			},
		}

		var str String
		if err := s.GetObject("https://example.com/ns#name").UnmarshalTo(&str); err != nil {
			t.Errorf("expected no error but got %s", err.Error())
		}
		if str != "John Doe" {
			t.Errorf("expected %s but got %s", "John Doe", str)
		}
	})
}

func TestUnmarshalTo(t *testing.T) {
	t.Run("unmarshal to", func(t *testing.T) {
		s := UnknownNode{
			"https://example.com/ns#name": map[string]any{
				"@value": "John Doe",
			},
		}

		var doc struct {
			Name String `json:"https://example.com/ns#name"`
		}
		if err := s.UnmarshalTo(&doc); err != nil {
			t.Errorf("expected no error but got %s", err.Error())
		}
		if string(doc.Name) != "John Doe" {
			t.Errorf("expected %s but got %s", "John Doe", doc.Name)
		}
	})
}
