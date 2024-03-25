package prettyld_test

import (
	"testing"

	"github.com/getfederated/prettyld"
)

func TestWithContext(t *testing.T) {
	t.Run("with context", func(t *testing.T) {
		source := []map[string]any{{
			"@id": "https://example.com/1",
			"https://example.com/ns#name": []map[string]any{{
				"@value": "John Doe",
			}},
		}}

		b, err := prettyld.WithContext(map[string]any{
			"ex":   "https://example.com/ns#",
			"name": "ex:name",
		}).MarshalCompactJSONLD(source, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		expected := `{"@context":{"ex":"https://example.com/ns#","name":"ex:name"},"@id":"https://example.com/1","name":"John Doe"}`
		if string(b) != expected {
			t.Errorf("expected `%s` but got `%s`", expected, string(b))
		}
	})
}
