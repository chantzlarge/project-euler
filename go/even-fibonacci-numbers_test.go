package euler

import "testing"

func TestEvenFibonacciNumbers(t *testing.T) {
	got := EvenFibonacciNumbers()
	expected := 4613732
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
	t.Logf("got:\t%v", got)
}
