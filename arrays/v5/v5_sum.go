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
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}
	return sums
}
