package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(sets ...[]int) []int {
	var sums []int
	for _, numbers := range sets {
		sums = append(sums, Sum(numbers))
	} // concat
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

// func SumAllTails(sets ...[]int) []int {
// 	lengthOfNumbers := len(sets)
// 	sums := make([]int, lengthOfNumbers)
// 	for index, numbers := range sets {
// 		length := len(numbers)
// 		sum := 0
// 		for i := 1; i < length; i++ {
// 			sum += numbers[i]
// 		}
// 		sums[index] = sum

// 	}
// 	return sums
// }

// func SumAll(sets ...[]int) []int {
// 	lengthOfNumbers := len(sets)
// 	sums := make([]int, lengthOfNumbers)
// 	for i, numbers := range sets {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }
