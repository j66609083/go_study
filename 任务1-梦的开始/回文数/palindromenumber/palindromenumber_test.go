package main

import "testing"

func TestIsPalindromeCase1(t *testing.T) {
	x := 121
	want := true
	got := isPalindrome(x)
	if got != want {
		t.Errorf("isPalindrome(%d) = %v, 预期结果: %v", x, got, want)
	}
}

func TestIsPalindromeCase2(t *testing.T) {
	x := -121
	want := false
	got := isPalindrome(x)
	if got != want {
		t.Errorf("isPalindrome(%d) = %v, 预期结果: %v", x, got, want)
	}
}

func TestIsPalindromeCase3(t *testing.T) {
	x := 10
	want := false
	got := isPalindrome(x)
	if got != want {
		t.Errorf("isPalindrome(%d) = %v, 预期结果: %v", x, got, want)
	}
}
