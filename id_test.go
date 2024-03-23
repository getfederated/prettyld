package prettyld_test

import (
	"encoding/json"
	"testing"

	"github.com/getfederated/prettyld"
)

func TestParseSingleNdoe(t *testing.T) {
	payload := `{"@id": "https://example.com"}`
	expected := "https://example.com"

	var id prettyld.ID

	err := json.Unmarshal([]byte(payload), &id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if string(id) != expected {
		t.Errorf("expected %s but got %s", expected, string(id))
	}
}

func TestParseSingleNodeInArray(t *testing.T) {
	payload := `[{"@id": "https://example.com"}]`
	expected := "https://example.com"

	var id prettyld.ID

	err := json.Unmarshal([]byte(payload), &id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if string(id) != expected {
		t.Errorf("expected %s but got %s", expected, string(id))
	}
}
