package euler

// EvenFibonacciNumbers finds the sum of all even Fibonacci numbers up to 4
// million.
func EvenFibonacciNumbers() int {
	prv := 0
	nxt := 1
	sum := 0
	for nxt < 4000000 {
		if nxt%2 == 0 {
			sum += nxt
		}
		tmp := nxt
		nxt += prv
		prv = tmp
	}
	return sum
}
