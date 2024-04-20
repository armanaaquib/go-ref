package arrays

import "testing"

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
