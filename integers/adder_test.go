package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add 10 and 9", func(t *testing.T) {
		sum := Add(10, 9)
		expected := 19

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}

	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
