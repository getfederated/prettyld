package prettyld

import "testing"

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

	p, err := Parse(b, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	var person Person
	err = p.Unmarshal(&person)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := "https://example.com"
	if string(person.ID) != expected {
		t.Errorf("expected %s but got %s", expected, string(person.ID))
	}
}
