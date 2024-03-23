package prettyld_test

import (
	"slices"
	"testing"

	"github.com/getfederated/prettyld"
)

type Person struct {
	ID   string                     `json:"@id"`
	Name prettyld.ValueNode[string] `json:"https://example.com/ns#name"`
}

func TestParseAndUnmarshal(t *testing.T) {
	b := []byte(`{
		"@context": {
			"ex": "https://example.com/ns#",
			"name": "ex:name"
		},
		"@id": "https://example.com",
		"name": "Alice"
	}`)

	type Person struct {
		ID   string                     `json:"@id"`
		Name prettyld.ValueNode[string] `json:"https://example.com/ns#name"`
	}

	p, err := prettyld.Parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	var person Person
	err = p.UnmarshalTo(&person)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := "https://example.com"
	if string(person.ID) != expected {
		t.Errorf("expected %s but got %s", expected, string(person.ID))
	}
}

func TestItems(t *testing.T) {
	b := []byte(`[
		{
			"@context": {
				"ex": "https://example.com/ns#",
				"name": "ex:name"
			},
			"@id": "https://example.com/1",
			"name": "Alice"
		},
		{
			"@context": {
				"ex": "https://example.com/ns#",
				"name": "ex:name"
			},
			"@id": "https://example.com/2",
			"name": "Bob"
		}
	]`)

	p, err := prettyld.Parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := []Person{{
		ID:   "https://example.com/1",
		Name: prettyld.ValueNode[string]{Value: "Alice"},
	}, {
		ID:   "https://example.com/2",
		Name: prettyld.ValueNode[string]{Value: "Bob"},
	}}

	actual := []Person{}

	for item := range p.Items() {
		var person Person

		err := item.UnmarshalTo(&person)
		if err != nil {
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

func TestNesting(t *testing.T) {
	b := []byte(`{
		"@context": {
			"ex": "https://example.com/ns#",
			"name": "ex:name",
			"friend": {
				"@id": "ex:friend",
				"@type": "@id"
			}
		},
		"@id": "https://example.com/1",
		"name": "Alice",
		"friend": {
			"@id": "https://example.com/2",
			"name": "Bob"
		}
	}`)

	type Person struct {
		ID     string                     `json:"@id"`
		Name   prettyld.ValueNode[string] `json:"https://example.com/ns#name"`
		Friend prettyld.LDNodesList       `json:"https://example.com/ns#friend"`
	}

	p, err := prettyld.Parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	var person Person
	err = p.UnmarshalTo(&person)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := "https://example.com/1"
	if person.ID != expected {
		t.Errorf("expected %s but got %s", expected, person.ID)
	}

	expected = "Alice"
	if person.Name.Value != expected {
		t.Errorf("expected %s but got %s", expected, person.Name.Value)
	}

	expected = "https://example.com/2"
	for item := range person.Friend.Items() {
		var friend Person

		err := item.UnmarshalTo(&friend)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if friend.ID != expected {
			t.Errorf("expected %s but got %s", expected, friend.ID)
		}
	}
}
