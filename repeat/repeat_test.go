package repeat

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
