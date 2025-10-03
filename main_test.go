package vskvuserHelloWorld

import "testing"

func TestHelloWorld(t *testing.T) {
	var expected = "Hello World";
	var got = GetHelloWorldStr();

	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
	}
}