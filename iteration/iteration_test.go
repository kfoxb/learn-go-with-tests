package iteration

import "testing"

func TestRepeated(t *testing.T) {
	actual := Repeated("j")
	expected := "jjjjj"
	if actual != expected {
		t.Errorf("expected '%s' but received '%s'", expected, actual)
	}
}
