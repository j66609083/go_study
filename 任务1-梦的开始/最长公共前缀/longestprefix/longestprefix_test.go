package main

import "testing"

func TestLongestCommonPrefixCase1(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	want := "fl"
	got := longestCommonPrefix(strs)
	if got != want {
		t.Errorf("longestCommonPrefix(%#v) = %q, 预期结果: %q", strs, got, want)
	}
}

func TestLongestCommonPrefixCase2(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}
	want := ""
	got := longestCommonPrefix(strs)
	if got != want {
		t.Errorf("longestCommonPrefix(%#v) = %q, 预期结果: %q", strs, got, want)
	}
}
