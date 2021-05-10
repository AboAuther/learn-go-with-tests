package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collections of any size", func(t *testing.T) {
		nums := []int{1, 23, 4}
		got := Sum(nums)
		want := 28
		if got != want {
			t.Errorf("got %d,want %d,%v", got, want, nums)
		}
	})
}
func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})
	t.Run("safety sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 6})
		want := []int{0, 6}
		checkSums(t, got, want)
	})
}
