package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, a := range tab {
		if f(a) == true {
			count++
		}
	}
	return count
}
