package sum_slice

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, n := range numbersToSum {
		sums = append(sums, Sum(n))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, n := range numbersToSum {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			tails := n[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
