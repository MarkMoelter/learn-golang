package walk

import "testing"

func TestWalk(t *testing.T) {
	t.Run("check the number of function calls", func(t *testing.T) {
		const lenDesired = 1

		expected := "Chris"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != lenDesired {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), lenDesired)
		}

		if got[0] != expected {
			t.Errorf("got %q, want %q", got[0], expected)
		}
	})
}
