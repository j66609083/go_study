package main

import (
	"reflect"
	"testing"
)

func TestPlusOneCase1(t *testing.T) {
	digits := []int{1, 2, 3}
	want := []int{1, 2, 4}
	got := plusOne(digits)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("plusOne(%#v) = %q, 预期结果: %q", digits, got, want)
	}
}

func TestPlusOneCase2(t *testing.T) {
	digits := []int{4, 3, 2, 1}
	want := []int{4, 3, 2, 2}
	got := plusOne(digits)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("plusOne(%#v) = %q, 预期结果: %q", digits, got, want)
	}
}

func TestPlusOneCase3(t *testing.T) {
	digits := []int{9}
	want := []int{1, 0}
	got := plusOne(digits)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("plusOne(%#v) = %q, 预期结果: %q", digits, got, want)
	}
}
