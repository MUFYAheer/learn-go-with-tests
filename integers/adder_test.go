package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCurrentMessage := func(t testing.TB, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("add 2 and 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCurrentMessage(t, sum, expected)
	})

	t.Run("add 2 and 3", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		assertCurrentMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
