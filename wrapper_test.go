package prettyld_test

import (
	"encoding/json"
	"testing"

	"github.com/getfederated/prettyld"
)

type Nilable[T any] struct {
	hasValue bool
	value    T
}

func Just[T any](v T) Nilable[T] {
	return Nilable[T]{
		hasValue: true,
		value:    v,
	}
}

func Nil[T any]() Nilable[T] {
	var t T
	return Nilable[T]{
		hasValue: false,
		value:    t,
	}
}

func (n Nilable[T]) HasValue() bool {
	return n.hasValue
}

func (n Nilable[T]) Value() (T, bool) {
	if !n.hasValue {
		return n.value, false
	}

	return n.value, true
}

func (n Nilable[T]) ValueOrDefault(d T) T {
	if !n.hasValue {
		return d
	}

	return n.value
}

func (n Nilable[T]) AssertValue() T {
	if !n.hasValue {
		panic("null dereference error")
	}
	return n.value
}

func (n Nilable[T]) MarshalJSON() ([]byte, error) {
	if !n.hasValue {
		return []byte("null"), nil
	}

	return json.Marshal(n.value)
}

func (n *Nilable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.hasValue = false
		var t T
		n.value = t
		return nil
	}

	var value T
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	n.hasValue = true
	n.value = value
	return nil
}

func Then[T any, V any](n Nilable[T], fn func(T) Nilable[V]) Nilable[V] {
	if !n.hasValue {
		return Nil[V]()
	}

	return fn(n.value)
}

func TestUnmarshalNilable(t *testing.T) {
	t.Run(`If a "string" value node has been set, then nilable should be able to pick it up`, func(t *testing.T) {
		type Object struct {
			Name Nilable[prettyld.String] `json:"https://example.com/ns#name"`
		}

		source := `{"https://example.com/ns#name": "John Doe"}`

		var obj Object
		err := prettyld.Unmarshal([]byte(source), &obj, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if !obj.Name.HasValue() {
			t.Error("expected Name to have a value")
			t.FailNow()
		}

		if name, ok := obj.Name.Value(); !ok || name != "John Doe" {
			t.Error("expected Name to be John Doe")
			t.FailNow()
		}
	})

	t.Run(`If a "string" value node has not been set, then nilable should be set to nil`, func(t *testing.T) {
		type Object struct {
			Name Nilable[prettyld.String] `json:"https://example.com/ns#name"`
		}

		source := `{"https://example.com/ns#something": true}`

		var obj Object
		err := prettyld.Unmarshal([]byte(source), &obj, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if obj.Name.HasValue() {
			t.Error("expected Name to not have a value")
			t.FailNow()
		}
	})

	t.Run(`If a node has not been set, then nilable should be set to nil`, func(t *testing.T) {
		type Object struct {
			Name Nilable[prettyld.ID] `json:"https://example.com/ns#friend"`
		}

		source := `{"https://example.com/ns#friend":{"@id":"https://example.com/users/johndoe"}}`

		var obj Object
		err := prettyld.Unmarshal([]byte(source), &obj, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if !obj.Name.HasValue() {
			t.Error("expected Name to have a value")
			t.FailNow()
		}

		expected := "https://example.com/users/johndoe"
		if obj.Name.ValueOrDefault(prettyld.ID("")) != prettyld.ID(expected) {
			t.Errorf("expected Name to be %s", expected)
			t.FailNow()
		}
	})
}

func TestMarshalNilable(t *testing.T) {
	t.Run("We should be able to marshal a nilable value that is not nil", func(t *testing.T) {
		type Object struct {
			Name Nilable[prettyld.String] `json:"https://example.com/ns#name"`
		}

		obj := Object{
			Name: Just(prettyld.String("John Doe")),
		}

		b, err := prettyld.WithContext(nil).MarshalCompactJSONLD(obj, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		expected := `{"@context":null,"https://example.com/ns#name":"John Doe"}`
		if string(b) != expected {
			t.Errorf("expected %s, got %s", expected, string(b))
			t.FailNow()
		}
	})
	t.Run("Marshaling a null should yield nothing", func(t *testing.T) {
		type Object struct {
			Name Nilable[prettyld.String] `json:"https://example.com/ns#name"`
		}

		obj := Object{
			Name: Nil[prettyld.String](),
		}

		b, err := prettyld.WithContext(nil).MarshalCompactJSONLD(obj, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		expected := `{}`
		if string(b) != expected {
			t.Errorf("expected %s, got %s", expected, string(b))
			t.FailNow()
		}
	})
}
