package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want)
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 5}, []int{4, 7})
	want := []int{8, 11}

	assertCorrectMessage(t, got, want)
}

// Sum of all items except first one
func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{2, 9}, []int{3, 4, 7})
	want := []int{9, 11}

	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %q got %q", want, got)
	}
}
