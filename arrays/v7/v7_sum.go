package main

// Sum 计算切片总和
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// SumAllTails 计算给定切片集合的除第一个数字以外的所有数字的和。
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, v := range numbersToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
