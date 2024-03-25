package prettyld

import (
	"slices"
	"testing"
)

type Person struct {
	ID   string            `json:"@id"`
	Name ValueNode[string] `json:"https://example.com/ns#name"`
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
		ID   string            `json:"@id"`
		Name ValueNode[string] `json:"https://example.com/ns#name"`
	}

	p, err := parse(b, nil)
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

	p, err := parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := []Person{{
		ID:   "https://example.com/1",
		Name: ValueNode[string]{Value: "Alice"},
	}, {
		ID:   "https://example.com/2",
		Name: ValueNode[string]{Value: "Bob"},
	}}

	actual := []Person{}

	for _, item := range p {
		var person Person

		nodes, err := parse(item, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		nodes.UnmarshalTo(&person)

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
		ID     string            `json:"@id"`
		Name   ValueNode[string] `json:"https://example.com/ns#name"`
		Friend LDNodesList       `json:"https://example.com/ns#friend"`
	}

	p, err := parse(b, nil)
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
	for _, item := range person.Friend {
		var friend Person

		nodes, err := parse(item, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		err = nodes.UnmarshalTo(&friend)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if friend.ID != expected {
			t.Errorf("expected %s but got %s", expected, friend.ID)
		}
	}
}

func TestParseIsMSA(t *testing.T) {
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

	p, err := parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(p) <= 0 {
		t.Error("expected more than 0 nodes")
	}
}

func TestIterate(t *testing.T) {
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

	p, err := parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := []Person{{
		ID:   "https://example.com/1",
		Name: ValueNode[string]{Value: "Alice"},
	}, {
		ID:   "https://example.com/2",
		Name: ValueNode[string]{Value: "Bob"},
	}}

	actual := []Person{}

	for v := range p.Iterate() {
		var person Person

		nodes, err := parse(v, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		nodes.UnmarshalTo(&person)

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

	if len(actual) != len(expected) {
		t.Errorf("expected 2 but got %d", len(actual))
	}

	for i, person := range actual {
		if person.ID != expected[i].ID {
			t.Errorf("expected %s but got %s", expected[i].ID, person.ID)
		}
		if person.Name.Value != expected[i].Name.Value {
			t.Errorf("expected %s but got %s", expected[i].Name.Value, person.Name.Value)
		}
	}
}

func TestUnmarshalToSlice(t *testing.T) {
	b := []byte(`{
		"@context": {
			"ex": "https://example.com/ns#",
			"name": "ex:name"
		},
		"@id": "https://example.com/1",
		"name": "Alice"
	}`)

	p, err := parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	var person []Person
	err = p.UnmarshalTo(&person)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(person) != 1 {
		t.Errorf("expected 1 but got %d", len(person))
	}

	expected := "https://example.com/1"
	if person[0].ID != expected {
		t.Errorf("expected %s but got %s", expected, person[0].ID)
	}

	expected = "Alice"
	if person[0].Name.Value != expected {
		t.Errorf("expected %s but got %s", expected, person[0].Name.Value)
	}
}

func TestUnmarshalToString(t *testing.T) {
	p := LDNodesList{{"@value": "Alice"}}

	var str String
	err := p.UnmarshalTo(&str)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := "Alice"
	if str != String(expected) {
		t.Errorf("expected %s but got %s", expected, str)
	}
}
