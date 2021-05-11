package main

// Sum 计算切片总和
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// SumAll 计算传入的每个切片的各自总和。
func SumAll(numberToSum ...[]int) []int {
	lengthOfNumbers := len(numberToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numberToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
