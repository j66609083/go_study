package main

import "testing"

func TestIsValidCase1(t *testing.T) {
	s := "()"
	want := true
	got := isValid(s)
	if got != want {
		t.Errorf("isValid(%q) = %v, 预期结果: %v", s, got, want)
	}
}

func TestIsValidCase2(t *testing.T) {
	s := "()[]{}"
	want := true
	got := isValid(s)
	if got != want {
		t.Errorf("isValid(%q) = %v, 预期结果: %v", s, got, want)
	}
}

func TestIsValidCase3(t *testing.T) {
	s := "(]"
	want := false
	got := isValid(s)
	if got != want {
		t.Errorf("isValid(%q) = %v, 预期结果: %v", s, got, want)
	}
}

func TestIsValidCase4(t *testing.T) {
	s := "([])"
	want := true
	got := isValid(s)
	if got != want {
		t.Errorf("isValid(%q) = %v, 预期结果: %v", s, got, want)
	}
}

func TestIsValidCase5(t *testing.T) {
	s := "([)]"
	want := false
	got := isValid(s)
	if got != want {
		t.Errorf("isValid(%q) = %v, 预期结果: %v", s, got, want)
	}
}
