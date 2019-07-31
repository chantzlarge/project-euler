package euler

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	expected := 6857
	got := LargestPrimeFactor()
	if got != 6857 {
		t.Errorf("expected %v, got %v", expected, got)
	}
	t.Logf("got:\t%v", got)
}
