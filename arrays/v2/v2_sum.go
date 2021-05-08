package main

func Sum(nums [5]int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}
