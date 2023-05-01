package piscine

func ActiveBits(n int) int {
	var a []int
	var count int
	for i := 0; i < 8; i++ {
		if n%2 == 1 {
			count++
		}
		a = append(a, n%2)

		n /= 2
	}
	return count
}
