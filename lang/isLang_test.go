package lo

import "testing"

func TestIsNumberWithString(t *testing.T) {
	result := IsNumber("hello");
	if result {
		t.Error("hello isn't Number");
		t.Fail();
	}
}

func TestIsNumberForNumber(t *testing.T) {
	shouldBeNumber := IsNumber("23456");
	if !shouldBeNumber {
		t.Error("23456 is a Number");
		t.Fail();
	}
}
