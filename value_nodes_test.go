package prettyld

import "testing"

func TestBoolMarshalJSON(t *testing.T) {
	t.Run("marshal bool", func(t *testing.T) {
		s := Bool(false)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":false}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestBoolUnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Bool
		err := s.UnmarshalJSON([]byte(`{"@value":true}`))
		if err != nil {
			t.Error(err)
		}
		expected := Bool(true)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestStringMarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := String("foo")
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":"foo"}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestStringUnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s String
		err := s.UnmarshalJSON([]byte(`{"@value":"foo"}`))
		if err != nil {
			t.Error(err)
		}
		expected := String("foo")
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestIntMarshalJSON(t *testing.T) {
	t.Run("marshal int", func(t *testing.T) {
		s := Int(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestIntUnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Int
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Int(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestInt8MarshalJSON(t *testing.T) {
	t.Run("marshal int8", func(t *testing.T) {
		s := Int8(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestInt8UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Int8
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Int8(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestInt16MarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Int16(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestInt16UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Int16
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Int16(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestInt32MarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Int32(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestInt32UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Int32
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Int32(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestInt64MarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Int64(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestInt64UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Int64
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Int64(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestUintMarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Uint(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestUintUnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Uint
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Uint(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestUint8MarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Uint8(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestUint8UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Uint8
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Uint8(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestUint16MarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Uint16(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestUin16UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Uint16
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Uint16(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestUint32MarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Uint32(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestUin32UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Uint32
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Uint32(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}

func TestUint64MarshalJSON(t *testing.T) {
	t.Run("marshal string", func(t *testing.T) {
		s := Uint64(42)
		b, err := s.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		expected := `{"@value":42}`
		if string(b) != expected {
			t.Errorf("expected %s but got %s", expected, string(b))
		}
	})
}

func TestUin64UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal bool", func(t *testing.T) {
		var s Uint64
		err := s.UnmarshalJSON([]byte(`{"@value":42}`))
		if err != nil {
			t.Error(err)
		}
		expected := Uint64(42)
		if s != expected {
			t.Errorf("expected %v but got %v", expected, s)
		}
	})
}
