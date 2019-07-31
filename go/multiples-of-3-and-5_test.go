package euler

import "testing"

func TestMultiplesOf3And5(t *testing.T) {
	g := MultiplesOf3And5(1000)
	e := 233168
	if g != e {
		t.Errorf("expected %v, got %v", e, g)
	}
	t.Logf("got:\t%v", g)
}
