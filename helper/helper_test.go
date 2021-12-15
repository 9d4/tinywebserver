package helper

import "testing"

func TestIsEmptyStringFunc(t *testing.T) {
	input := ""
	expected := true

	result := IsEmptyString(input)
	if result != expected {
		t.Errorf("Expected %t, Input: %s, Got: %v", expected, input, result)
	}
	
	input = "asdasd"
	expected = false 

	result = IsEmptyString(input)
	if result != expected {
		t.Errorf("Expected %t, Input: %s, Got: %v", expected, input, result)
	}
}
