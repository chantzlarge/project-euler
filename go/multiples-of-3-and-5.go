package euler

// MultiplesOf3And5 finds the sum of all the multiples of 3 or 5 below 1000.
func MultiplesOf3And5() int {
	sum := 0
	for n := 0; n < 1000; n++ {
		if n%3 == 0 || n%5 == 0 {
			sum += n
		}
	}
	return sum
}
