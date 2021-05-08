package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6
		if got != want {
			t.Errorf("got:%d,want:%d,%v", got, want, nums)
		}
	})
}
