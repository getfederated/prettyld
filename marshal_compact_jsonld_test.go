package prettyld

import (
	"testing"
)

func TestIsValidObject(t *testing.T) {
	t.Run("valid for empty struct", func(t *testing.T) {
		if !isValidObject(struct{}{}) {
			t.Error("isValidObject should return true for a struct")
		}
	})

	t.Run("valid for empty pointer to struct", func(t *testing.T) {
		if !isValidObject(&struct{}{}) {
			t.Error("isValidObject should return true for a pointer to a struct")
		}
	})

	t.Run("invalid for int", func(t *testing.T) {
		if isValidObject(1) {
			t.Error("isValidObject should return false for an int")
		}
	})

	t.Run("valid for slice of any", func(t *testing.T) {
		if !isValidObject([]any{}) {
			t.Error("isValidObject should return true for an empty slice")
		}
	})

	t.Run("valid for slice of structs", func(t *testing.T) {
		if !isValidObject([]struct{ Foo string }{{Foo: "bar"}}) {
			t.Error("isValidObject should return true for a slice of structs")
		}
	})

	t.Run("valid for slice of struct typed as []any", func(t *testing.T) {
		if !isValidObject([]any{struct{ Foo string }{Foo: "bar"}}) {
			t.Error("isValidObject should return true for a slice of structs")
		}
	})

	t.Run("invalid for a slice containing a mixture of struct and non-struct", func(t *testing.T) {
		if isValidObject([]any{struct{ Foo string }{Foo: "bar"}, 1}) {
			t.Error("isValidObject should return false for a slice containing a mixture of struct and non-struct")
		}
	})
}

func TestIsMSA(t *testing.T) {
	t.Run("something not MSA should yield false", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			if isMSA("foo") {
				t.Error("isMSA should return false for a string")
			}
		})
		t.Run("number", func(t *testing.T) {
			if isMSA(10) {
				t.Error("isMSA should return false for a string")
			}
		})
	})

	t.Run("something that is MSA should yield true", func(t *testing.T) {
		t.Run("map[string]struct{}", func(t *testing.T) {
			if !isMSA(map[string]struct{}{}) {
				t.Error("isMSA should return true for a map[string]struct{}")
			}
		})
		t.Run("map[string]struct{} pointer", func(t *testing.T) {
			if !isMSA(&map[string]struct{}{}) {
				t.Error("isMSA should return true for a pointer to a map[string]struct{}")
			}
		})
		t.Run("map[int]struct{} pointer", func(t *testing.T) {
			if isMSA(&map[int]struct{}{}) {
				t.Error("isMSA should return false for a pointer to a map[int]struct{}")
			}
		})
	})
}

func TestMarshalCompactJSONLD(t *testing.T) {
	t.Run("should yield an error when attempting to marshal something that isn't a struct", func(t *testing.T) {
		_, err := marshalCompactJSONLD(1, nil, nil)
		if err == nil {
			t.Error("marshalCompactJSONLD should return an error for an int")
		}
	})

	t.Run("should yield an error when attempting to marshal a slice of ints", func(t *testing.T) {
		_, err := marshalCompactJSONLD([]any{1}, nil, nil)
		if err == nil {
			t.Error("marshalCompactJSONLD should return an error for a slice of ints")
		}
	})

	t.Run("should marshal a struct with no valid predicate IRIs into an empty JSON object", func(t *testing.T) {
		b, err := marshalCompactJSONLD(struct{ Foo string }{Foo: "bar"}, nil, nil)
		if err != nil {
			t.Error("marshalCompactJSONLD should not return an error for a struct")
		}
		if string(b) != `{}` {
			t.Error("marshalCompactJSONLD should return an empty object for a struct")
		}
	})

	t.Run("should marshal a struct with a well-defined predicate IRI into a JSON object with a single field", func(t *testing.T) {
		b, err := marshalCompactJSONLD(&struct {
			Foo string `json:"https://example.com/foo"`
		}{Foo: "bar"}, nil, nil)
		if err != nil {
			t.Error("marshalCompactJSONLD should not return an error for a pointer to a struct")
		}
		expected := `{"@context":null,"https://example.com/foo":"bar"}`
		if string(b) != expected {
			t.Errorf("expected `%s` but got `%s`", expected, string(b))
		}
	})

	t.Run("should marshal a slice of structs with no valid predicate IRIs into an empty JSON object", func(t *testing.T) {
		b, err := marshalCompactJSONLD([]struct{ Foo string }{{Foo: "bar"}}, nil, nil)
		if err != nil {
			t.Error("marshalCompactJSONLD should not return an error for a slice of structs")
		}
		if string(b) != `{}` {
			t.Error("marshalCompactJSONLD should return an empty object for a slice of structs that lack proper fields")
		}
	})

	t.Run("should mrshal a slice of structs with a well-defined predicate IRI into a JSON object with a single field", func(t *testing.T) {
		b, err := marshalCompactJSONLD([]struct {
			Foo string `json:"https://example.com/foo"`
		}{{Foo: "bar"}}, nil, nil)
		if err != nil {
			t.Error("marshalCompactJSONLD should not return an error for a pointer to a struct")
		}
		expected := `{"@context":null,"https://example.com/foo":"bar"}`
		if string(b) != `{"@context":null,"https://example.com/foo":"bar"}` {
			t.Errorf("expected `%s` but got `%s`", expected, string(b))
		}
	})

	t.Run("should marshal a slice of structs with no valid predicate IRIs into an empty JSON object", func(t *testing.T) {
		b, err := marshalCompactJSONLD(&[]struct{ Foo string }{{Foo: "bar"}}, nil, nil)
		if err != nil {
			t.Error("marshalCompactJSONLD should not return an error for a slice of structs")
		}
		if string(b) != `{}` {
			t.Error("marshalCompactJSONLD should return an empty object for a slice of structs that lack proper fields")
		}
	})

	t.Run("should marshal a slice of structs with a well-defined predicate IRI into a JSON object with a single field", func(t *testing.T) {
		b, err := marshalCompactJSONLD(&[]struct {
			Foo string `json:"https://example.com/foo"`
		}{{Foo: "bar"}}, nil, nil)
		if err != nil {
			t.Error("marshalCompactJSONLD should not return an error for a pointer to a struct")
		}
		expected := `{"@context":null,"https://example.com/foo":"bar"}`
		if string(b) != expected {
			t.Errorf("expected `%s` but got `%s`", expected, string(b))
		}
	})
}

func TestMarshalExpanded(t *testing.T) {
	value := []map[string]any{{
		"https://example.com/ns#name": []map[string]any{{
			"@value": "John Doe",
		}},
	}}

	b, err := marshalCompactJSONLD(value, nil, nil)
	if err != nil {
		t.Error("marshalCompactJSONLD should not return an error for a struct")
		t.Errorf("error: %s", err.Error())
		t.FailNow()
	}

	expected := `{"@context":null,"https://example.com/ns#name":"John Doe"}`
	if string(b) != expected {
		t.Errorf("expected `%s` but got `%s`", expected, string(b))
	}
}

func TestMarshalWithContext(t *testing.T) {
	type someStruct struct {
		Name string `json:"https://example.com/ns#name"`
	}

	value := someStruct{
		Name: "John Doe",
	}

	b, err := marshalCompactJSONLD(value, map[string]any{
		"@context": map[string]any{
			"ex":   "https://example.com/ns#",
			"name": "ex:name",
		},
	}, nil)
	if err != nil {
		t.Error("marshalCompactJSONLD should not return an error for a struct")
		t.Errorf("error: %s", err.Error())
		t.FailNow()
	}

	expected := `{"@context":{"ex":"https://example.com/ns#","name":"ex:name"},"name":"John Doe"}`
	if string(b) != expected {
		t.Errorf("expected `%s` but got `%s`", expected, string(b))
	}
}
