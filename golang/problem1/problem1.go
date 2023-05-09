package problem1

func Mul3of5(N int) int {
	var sum int
	for i := 1; i < N; i++ {
		isDivisible := (i%3 == 0 || i%5 == 0)
		if isDivisible {
			sum += i

		}
	}
	return sum
}
