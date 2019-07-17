package euler

import "testing"

func TestMultiplesOf3And5(t *testing.T) {
	r := MultiplesOf3And5()
	e := 233168
	if r != e {
		t.Errorf("expected:\t%v\nfound:\t%v", e, r)
	}
	t.Logf("got:\t%v", r)
}
