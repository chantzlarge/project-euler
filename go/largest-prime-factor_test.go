package euler

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	e := 6857
	g := LargestPrimeFactor(600851475143)
	if g != e {
		t.Errorf("expected %v, got %v", e, g)
	}
	t.Logf("got:\t%v", g)
}
