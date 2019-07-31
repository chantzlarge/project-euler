package euler

// MultiplesOf3And5 finds the sum of all the multiples of 3 or 5 below 1000.
func MultiplesOf3And5(n int) int {
	t := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			t += i
		}
	}
	return t
}
