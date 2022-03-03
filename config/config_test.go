package config

import (
	"testing"
)

func TestHasFieldFunc(t *testing.T) {
	c := &configuration{HOST: ":9000"}

	if hasField(c, "HOST") != true {
		t.Errorf("Expected: %t, Key: HOST, Result: %v", true, hasField(c, "HOST"))
	}

	if hasField(c, "INVALIDKEY") == true {
		t.Errorf("Expected: %t, Key: INVALIDKEY, Result: %v", false, hasField(c, "INVALIDKEY"))
	}
}
