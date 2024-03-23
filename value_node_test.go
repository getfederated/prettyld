package prettyld

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalValueNode(t *testing.T) {
	payload := `{"@value": "foobar"}`
	expected := "foobar"

	var vn ValueNode[string]

	err := json.Unmarshal([]byte(payload), &vn)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if vn.Value != expected {
		t.Errorf("expected %s but got %s", expected, vn.Value)
	}
}

func TestUnmarshalValueNodeFromArray(t *testing.T) {
	payload := `[{"@value": "foobar"}]`
	expected := "foobar"

	var vn ValueNode[string]

	err := json.Unmarshal([]byte(payload), &vn)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if vn.Value != expected {
		t.Errorf("expected %s but got %s", expected, vn.Value)
	}
}
