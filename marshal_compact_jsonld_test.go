package prettyld

import "testing"

func TestIsValidObject(t *testing.T) {
	if !isValidObject(struct{}{}) {
		t.Error("isValidObject should return true for a struct")
	}

	if !isValidObject(&struct{}{}) {
		t.Error("isValidObject should return true for a pointer to a struct")
	}

	if isValidObject(1) {
		t.Error("isValidObject should return false for an int")
	}

	if !isValidObject([]any{}) {
		t.Error("isValidObject should return true for an empty slice")
	}

	if !isValidObject([]struct{ Foo string }{{Foo: "bar"}}) {
		t.Error("isValidObject should return true for a slice of structs")
	}

	if isValidObject([]any{struct{ Foo string }{Foo: "bar"}, 1}) {
		t.Error("isValidObject should return false for a slice containing a mixture of struct and non-struct")
	}
}

func TestMarshalCompactJSONLD(t *testing.T) {
	_, err := MarshalCompactJSONLD(1, nil)
	if err == nil {
		t.Error("MarshalCompactJSONLD should return an error for an int")
	}

	_, err = MarshalCompactJSONLD([]any{1}, nil)
	if err == nil {
		t.Error("MarshalCompactJSONLD should return an error for a slice of ints")
	}

	b, err := MarshalCompactJSONLD(struct{ Foo string }{Foo: "bar"}, nil)
	if err != nil {
		t.Error("MarshalCompactJSONLD should not return an error for a struct")
	}
	if string(b) != `{}` {
		t.Error("MarshalCompactJSONLD should return an empty object for a struct")
	}

	b, err = MarshalCompactJSONLD(&struct {
		Foo string `json:"https://example.com/foo"`
	}{Foo: "bar"}, nil)
	if err != nil {
		t.Error("MarshalCompactJSONLD should not return an error for a pointer to a struct")
	}
	expected := `{"@context":null,"https://example.com/foo":"bar"}`
	if string(b) != expected {
		t.Errorf("expected `%s` but got `%s`", expected, string(b))
	}

	b, err = MarshalCompactJSONLD([]struct{ Foo string }{{Foo: "bar"}}, nil)
	if err != nil {
		t.Error("MarshalCompactJSONLD should not return an error for a slice of structs")
	}
	if string(b) != `{}` {
		t.Error("MarshalCompactJSONLD should return an empty object for a slice of structs that lack proper fields")
	}

	b, err = MarshalCompactJSONLD([]struct {
		Foo string `json:"https://example.com/foo"`
	}{{Foo: "bar"}}, nil)
	if err != nil {
		t.Error("MarshalCompactJSONLD should not return an error for a pointer to a struct")
	}
	expected = `{"@context":null,"https://example.com/foo":"bar"}`
	if string(b) != `{"@context":null,"https://example.com/foo":"bar"}` {
		t.Errorf("expected `%s` but got `%s`", expected, string(b))
	}
}
