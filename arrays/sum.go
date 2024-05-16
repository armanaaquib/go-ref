package arrays

func Sum(numbers []int) int {
	add := func(acc, num int) int {
		return acc + num
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	sum := func(acc, numbers []int) []int {
		return append(acc, Sum(numbers))
	}
	return Reduce(numbersToSum, sum, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, numbers []int) []int {
		if len(numbers) == 0 {
			return append(acc, 0)
		} else {
			tail := numbers[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}
