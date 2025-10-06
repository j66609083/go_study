package main

import (
	"testing"
)

func TestSingleNumberCase1(t *testing.T) {
	nums := []int{2, 2, 1}
	want := 1
	got := singleNumber(nums)

	if got != want {
		t.Errorf("singleNumber(%#v) = %d, 预期结果: %d", nums, got, want)
	}
}

func TestSingleNumberCase2(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	want := 4
	got := singleNumber(nums)

	if got != want {
		t.Errorf("singleNumber(%#v) = %d, 预期结果: %d", nums, got, want)
	}
}

func TestSingleNumberCase3(t *testing.T) {
	nums := []int{1}
	want := 1
	got := singleNumber(nums)

	if got != want {
		t.Errorf("singleNumber(%#v) = %d, 预期结果: %d", nums, got, want)
	}
}
