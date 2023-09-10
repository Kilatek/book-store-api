package mapper

import (
	"testing"
)

func TestMapStructsWithJSONTags(t *testing.T) {

	type source struct {
		A string `json:"test"`
	}

	type dest struct {
		B string `json:"test"`
	}

	s := &source{
		A: "Value from source",
	}
	d := &dest{}

	if err := MapStructsWithJSONTags(s, d); err != nil {
		t.Errorf("MapStructsWithJSONTags() error = %v", err)
	}

	if d.B != s.A {
		t.Errorf("MapStructsWithJSONTags(): dest valua not equal source valua")
	}
}
