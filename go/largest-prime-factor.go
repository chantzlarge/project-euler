package euler

// LargestPrimeFactor returns the largest prime factor for the given value 600851475142.
func LargestPrimeFactor(n int) int {
	for i := 2; i < n; i++ {
		for n%i == 0 {
			n /= i
		}
	}

	return n
}
