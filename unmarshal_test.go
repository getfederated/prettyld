package prettyld_test

import (
	"testing"

	"github.com/getfederated/prettyld"
)

func TestUnmarshalString(t *testing.T) {
	type MyModel struct {
		ID   string          `json:"@id"`
		Type []string        `json:"@type"`
		Name prettyld.String `json:"https://example.com/ns#name"`
	}

	var j = `
		{
			"@context": {
				"ex": "https://example.com/ns#",
				"name": "ex:name"
			},
			"name": "Alice"
		}
	`

	var dest MyModel

	err := prettyld.Unmarshal(j, &dest, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if dest.Name != "Alice" {
		t.Errorf("expected `Alice` but got `%s`", dest.Name)
		t.Fail()
	}
}

func TestIsType(t *testing.T) {
	t.Run("should be true if node is of specific type", func(t *testing.T) {
		node := prettyld.UnknownNode{
			"@type": []string{"https://example.com/ns#Person"},
		}

		if !node.IsType("https://example.com/ns#Person") {
			t.Error("Expected true, but got false")
		}
	})
	t.Run("should be true if node is of specific type", func(t *testing.T) {
		node := prettyld.UnknownNode{
			"@type": []string{"https://example.com/ns#Animal"},
		}

		if node.IsType("https://example.com/ns#Person") {
			t.Error("Expected true, but got false")
		}
	})
}
