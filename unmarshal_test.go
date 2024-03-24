package prettyld_test

import (
	"fmt"
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

	err := prettyld.Unmarshal(j, &dest)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(dest)

	if dest.Name != "Alice" {
		t.Errorf("expected `Alice` but got `%s`", dest.Name)
		t.Fail()
	}

	t.FailNow()
}

func TestUnmarshalLDNodesList(t *testing.T) {
	type MyModel struct {
		ID      string                 `json:"@id"`
		Type    []string               `json:"@type"`
		Friends []prettyld.UnknownNode `json:"https://example.com/ns#friends"`
	}

	var j = `
		{
			"@context": {
				"ex": "https://example.com/ns#",
				"friends": {
					"@id": "ex:friends",
					"@type": "@id"
				},
				"name": "ex:name"
			},
			"@type": "https://example.com/Friendship",
			"friends": [
				{
					"@id": "https://example.com/1",
					"name": "Alice"
				},
				{
					"@id": "https://example.com/2",
					"name": "Bob"
				}
			]
		}
	`

	var dest MyModel

	err := prettyld.Unmarshal(j, &dest)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(dest.Friends) != 2 {
		t.Error("Expected 2 friends, but got ", len(dest.Friends))
	}
}
