package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collections of ant sizes", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6
		if got != want {
			t.Errorf("got :%d,want:%d,%v", got, want, nums)
		}
	})
}
func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{2, 3}, []int{3, 6})
	want := []int{2, 3, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v,want %v", got, want)
	}
}
