package euler

import "testing"

func TestEvenFibonacciNumbers(t *testing.T) {
	g := EvenFibonacciNumbers()
	e := 4613732
	if g != e {
		t.Errorf("expected %v, got %v", e, g)
	}
	t.Logf("got:\t%v", g)
}
