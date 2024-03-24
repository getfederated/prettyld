package prettyld

import "testing"

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

func TestMarshalCompactJSONLD(t *testing.T) {
	t.Run("should yield an error when attempting to marshal something that isn't a struct", func(t *testing.T) {
		_, err := MarshalCompactJSONLD(1, nil)
		if err == nil {
			t.Error("MarshalCompactJSONLD should return an error for an int")
		}
	})

	t.Run("should yield an error when attempting to marshal a slice of ints", func(t *testing.T) {
		_, err := MarshalCompactJSONLD([]any{1}, nil)
		if err == nil {
			t.Error("MarshalCompactJSONLD should return an error for a slice of ints")
		}
	})

	t.Run("should marshal a struct with no valid predicate IRIs into an empty JSON object", func(t *testing.T) {
		b, err := MarshalCompactJSONLD(struct{ Foo string }{Foo: "bar"}, nil)
		if err != nil {
			t.Error("MarshalCompactJSONLD should not return an error for a struct")
		}
		if string(b) != `{}` {
			t.Error("MarshalCompactJSONLD should return an empty object for a struct")
		}
	})

	t.Run("should marshal a struct with a well-defined predicate IRI into a JSON object with a single field", func(t *testing.T) {
		b, err := MarshalCompactJSONLD(&struct {
			Foo string `json:"https://example.com/foo"`
		}{Foo: "bar"}, nil)
		if err != nil {
			t.Error("MarshalCompactJSONLD should not return an error for a pointer to a struct")
		}
		expected := `{"@context":null,"https://example.com/foo":"bar"}`
		if string(b) != expected {
			t.Errorf("expected `%s` but got `%s`", expected, string(b))
		}
	})

	t.Run("should marshal a slice of structs with no valid predicate IRIs into an empty JSON object", func(t *testing.T) {
		b, err := MarshalCompactJSONLD([]struct{ Foo string }{{Foo: "bar"}}, nil)
		if err != nil {
			t.Error("MarshalCompactJSONLD should not return an error for a slice of structs")
		}
		if string(b) != `{}` {
			t.Error("MarshalCompactJSONLD should return an empty object for a slice of structs that lack proper fields")
		}
	})

	t.Run("should marshal a slice of structs with a well-defined predicate IRI into a JSON object with a single field", func(t *testing.T) {
		b, err := MarshalCompactJSONLD([]struct {
			Foo string `json:"https://example.com/foo"`
		}{{Foo: "bar"}}, nil)
		if err != nil {
			t.Error("MarshalCompactJSONLD should not return an error for a pointer to a struct")
		}
		expected := `{"@context":null,"https://example.com/foo":"bar"}`
		if string(b) != `{"@context":null,"https://example.com/foo":"bar"}` {
			t.Errorf("expected `%s` but got `%s`", expected, string(b))
		}
	})
}
