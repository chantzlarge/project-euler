package euler

// LargestPrimeFactor returns the largest prime factor for the given value 600851475142.
func LargestPrimeFactor() int {
	n := 600851475143

	for i := 2; i < n; i++ {
		for n%i == 0 {
			n /= i
		}
	}

	return n
}
