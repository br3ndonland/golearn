package main

import (
	"reflect"
	"testing"
)

func assertCorrectSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertDeepEqual(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertEachCorrectSum(t testing.TB, got, want []int) {
	// array comparison algorithm similar to reflect.DeepEqual()
	t.Helper()
	gotLen := len(got)
	wantLen := len(want)
	if gotLen != wantLen {
		t.Errorf("got slice is length %d but want slice is %d", gotLen, wantLen)
	}
	for i, value := range want {
		if value != got[i] {
			t.Errorf("got %d want %d", got[i], value)
		}
	}
}

func TestSumArrayOfFiveNumbers(t *testing.T) {
	t.Run("add 1-5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := SumArrayOfFiveNumbers(numbers)
		want := 15
		assertCorrectSum(t, got, want)
	})
	t.Run("add 6 instead of 5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 6}
		got := SumArrayOfFiveNumbers(numbers)
		want := 16
		assertCorrectSum(t, got, want)
	})
	t.Run("add an unordered array of numbers", func(t *testing.T) {
		numbers := [5]int{1, 10, 2, 20, 347}
		got := SumArrayOfFiveNumbers(numbers)
		want := 380
		assertCorrectSum(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 9}, []int{0, 9})
	want := []int{12, 9}
	assertEachCorrectSum(t, got, want)
	assertDeepEqual(t, got, want)
}
