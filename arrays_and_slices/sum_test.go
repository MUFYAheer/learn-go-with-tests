package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCurrentMessage := func(t testing.TB, got, want int, given []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d but want %d, given %v", got, want, given)
		}
	}

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		assertCurrentMessage(t, got, want, numbers[:])
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{3, 4, 1}, []int{}, []int{1})
	want := []int{5, 5, 0, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
