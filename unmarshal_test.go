package prettyld_test

import (
	"fmt"
	"slices"
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

	expected := []Person{{
		ID:   "https://example.com/1",
		Name: prettyld.ValueNode[string]{Value: "Alice"},
	}, {
		ID:   "https://example.com/2",
		Name: prettyld.ValueNode[string]{Value: "Bob"},
	}}

	actual := []Person{}

	for _, v := range dest.Friends {
		var person Person

		fmt.Println(v)

		nodes, err := prettyld.Parse(v, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if err := nodes.UnmarshalTo(&person); err != nil {
			t.Error(err)
			t.FailNow()
		}

		actual = append(actual, person)
	}

	slices.SortFunc(actual, func(i, j Person) int {
		if i.ID < j.ID {
			return -1
		} else if i.ID > j.ID {
			return 1
		}
		return 0
	})

	for i, person := range actual {
		if person.ID != expected[i].ID {
			t.Errorf("expected %s but got %s", expected[i].ID, person.ID)
		}
		if person.Name.Value != expected[i].Name.Value {
			t.Errorf("expected %s but got %s", expected[i].Name.Value, person.Name.Value)
		}
	}
}
