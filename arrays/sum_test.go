package main

import "testing"

func assertCorrectSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
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
