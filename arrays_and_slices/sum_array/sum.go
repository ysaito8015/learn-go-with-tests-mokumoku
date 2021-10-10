package sum

func Sum(array [5]int) int {
	sum := 0
	for _, n := range array {
		sum += n
	}
	return sum
}
