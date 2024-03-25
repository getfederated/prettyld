package prettyld

import (
	"encoding/json"
)

// Bool is a type that represents a bool value in a JSON-LD document.
//
// In reality, Bool is a convenience type that is intended for translating
// a `ValueNode` into a bool. This is useful when you know that a value is
// a bool and you want to avoid the overhead of working with a `ValueNode`.
type Bool bool

var _ json.Marshaler = Bool(true)
var _ json.Unmarshaler = (*Bool)(nil)

func (s Bool) MarshalJSON() ([]byte, error) {
	value := ValueNode[bool]{
		Value: bool(s),
	}
	return json.Marshal(value)
}

func (s *Bool) UnmarshalJSON(data []byte) error {
	v := ValueNode[bool]{Value: bool(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Bool(v.Value)
	return nil
}

// String is a type that represents a string value in a JSON-LD document.
//
// In reality, Bool is a convenience type that is intended for translating
// a `ValueNode` into a string. This is useful when you know that a value is
// a string and you want to avoid the overhead of working with a `ValueNode`.
type String string

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s String) MarshalJSON() ([]byte, error) {
	value := ValueNode[string]{
		Value: string(s),
	}
	return json.Marshal(value)
}

func (s *String) UnmarshalJSON(data []byte) error {
	v := ValueNode[string]{Value: string(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = String(v.Value)
	return nil
}

// Int is a type that represents a Go int value in a JSON-LD document.
//
// In reality, Int is a convenience type that is intended for translating
// a `ValueNode` into an int. This is useful when you know that a value is
// an int and you want to avoid the overhead of working with a `ValueNode`.
type Int int

var _ json.Marshaler = Int(42)
var _ json.Unmarshaler = (*Int)(nil)

func (s Int) MarshalJSON() ([]byte, error) {
	value := ValueNode[int]{
		Value: int(s),
	}
	return json.Marshal(value)
}

func (s *Int) UnmarshalJSON(data []byte) error {
	v := ValueNode[int]{Value: int(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Int(v.Value)
	return nil
}

// Int8 is a type that represents a Go int8 value in a JSON-LD document.
//
// In reality, Int8 is a convenience type that is intended for translating
// a `ValueNode` into an int8. This is useful when you know that a value is
// an int8 and you want to avoid the overhead of working with a `ValueNode`.
type Int8 int8

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Int8) MarshalJSON() ([]byte, error) {
	value := ValueNode[int8]{
		Value: int8(s),
	}
	return json.Marshal(value)
}

func (s *Int8) UnmarshalJSON(data []byte) error {
	v := ValueNode[int8]{Value: int8(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Int8(v.Value)
	return nil
}

// Int16 is a type that represents a Go int16 value in a JSON-LD document.
//
// In reality, In16 is a convenience type that is intended for translating
// a `ValueNode` into an int16. This is useful when you know that a value is
// an int16 and you want to avoid the overhead of working with a `ValueNode`.
type Int16 int16

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Int16) MarshalJSON() ([]byte, error) {
	value := ValueNode[int16]{
		Value: int16(s),
	}
	return json.Marshal(value)
}

func (s *Int16) UnmarshalJSON(data []byte) error {
	v := ValueNode[int16]{Value: int16(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Int16(v.Value)
	return nil
}

// Int32 is a type that represents a Go int32 value in a JSON-LD document.
//
// In reality, Int32 is a convenience type that is intended for translating
// a `ValueNode` into an int16. This is useful when you know that a value is
// an int32 and you want to avoid the overhead of working with a `ValueNode`.
type Int32 int32

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Int32) MarshalJSON() ([]byte, error) {
	value := ValueNode[int32]{
		Value: int32(s),
	}
	return json.Marshal(value)
}

func (s *Int32) UnmarshalJSON(data []byte) error {
	v := ValueNode[int32]{Value: int32(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Int32(v.Value)
	return nil
}

// Int64 is a type that represents a Go int64 value in a JSON-LD document.
//
// In reality, Int64 is a convenience type that is intended for translating
// a `ValueNode` into an int16. This is useful when you know that a value is
// an int64 and you want to avoid the overhead of working with a `ValueNode`.
type Int64 int64

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Int64) MarshalJSON() ([]byte, error) {
	value := ValueNode[int64]{
		Value: int64(s),
	}
	return json.Marshal(value)
}

func (s *Int64) UnmarshalJSON(data []byte) error {
	v := ValueNode[int64]{Value: int64(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Int64(v.Value)
	return nil
}

// Uint is a type that represents a Go uint value in a JSON-LD document.
//
// In reality, Uint is a convenience type that is intended for translating
// a `ValueNode` into an uint. This is useful when you know that a value is
// an uint and you want to avoid the overhead of working with a `ValueNode`.
type Uint uint

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Uint) MarshalJSON() ([]byte, error) {
	value := ValueNode[uint]{
		Value: uint(s),
	}
	return json.Marshal(value)
}

func (s *Uint) UnmarshalJSON(data []byte) error {
	v := ValueNode[uint]{Value: uint(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Uint(v.Value)
	return nil
}

// Uint8 is a type that represents a Go uint8 value in a JSON-LD document.
//
// In reality, Uint8 is a convenience type that is intended for translating
// a `ValueNode` into an uint8. This is useful when you know that a value is
// an uint8 and you want to avoid the overhead of working with a `ValueNode`.
type Uint8 uint8

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Uint8) MarshalJSON() ([]byte, error) {
	value := ValueNode[uint8]{
		Value: uint8(s),
	}
	return json.Marshal(value)
}

func (s *Uint8) UnmarshalJSON(data []byte) error {
	v := ValueNode[uint8]{Value: uint8(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Uint8(v.Value)
	return nil
}

// Uint16 is a type that represents a Go uint16 value in a JSON-LD document.
//
// In reality, Uint16 is a convenience type that is intended for translating
// a `ValueNode` into an uint16. This is useful when you know that a value is
// an uint16 and you want to avoid the overhead of working with a `ValueNode`.
type Uint16 uint16

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Uint16) MarshalJSON() ([]byte, error) {
	value := ValueNode[uint16]{
		Value: uint16(s),
	}
	return json.Marshal(value)
}

func (s *Uint16) UnmarshalJSON(data []byte) error {
	v := ValueNode[uint16]{Value: uint16(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Uint16(v.Value)
	return nil
}

// Uint32 is a type that represents a Go uint32 value in a JSON-LD document.
//
// In reality, Uint32 is a convenience type that is intended for translating
// a `ValueNode` into an uint32. This is useful when you know that a value is
// an uint32 and you want to avoid the overhead of working with a `ValueNode`.
type Uint32 uint32

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Uint32) MarshalJSON() ([]byte, error) {
	value := ValueNode[uint32]{
		Value: uint32(s),
	}
	return json.Marshal(value)
}

func (s *Uint32) UnmarshalJSON(data []byte) error {
	v := ValueNode[uint32]{Value: uint32(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Uint32(v.Value)
	return nil
}

// Uint64 is a type that represents a Go uint64 value in a JSON-LD document.
//
// In reality, Uint64 is a convenience type that is intended for translating
// a `ValueNode` into an uint64. This is useful when you know that a value is
// an uint64 and you want to avoid the overhead of working with a `ValueNode`.
type Uint64 uint64

var _ json.Marshaler = String("")
var _ json.Unmarshaler = (*String)(nil)

func (s Uint64) MarshalJSON() ([]byte, error) {
	value := ValueNode[uint64]{
		Value: uint64(s),
	}
	return json.Marshal(value)
}

func (s *Uint64) UnmarshalJSON(data []byte) error {
	v := ValueNode[uint64]{Value: uint64(*s)}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*s = Uint64(v.Value)
	return nil
}
