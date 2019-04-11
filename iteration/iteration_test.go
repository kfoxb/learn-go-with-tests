package iteration

import "testing"

func TestRepeated(t *testing.T) {
	actual := Repeat("j", 5)
	expected := "jjjjj"
	if actual != expected {
		t.Errorf("expected '%s' but received '%s'", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("j", 5)
	}
}
