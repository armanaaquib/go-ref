package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{3, 4, 5}

		got := Sum(numbers)
		want := 12

		if got != want {
			t.Errorf("got %d wanted %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d wanted, %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("non empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{10, 9, 1})
		want := []int{5, 10}

		checkSums(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{})
		want := []int{5, 0}

		checkSums(t, got, want)
	})

}
