package repeat_test

import (
	"fmt"
	"testing"

	"github.com/MarkMoelter/learn-golang/repeat"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' five times", func(t *testing.T) {
		repeated := repeat.Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat.Repeat("a")
	}
}

func ExampleRepeat() {
	repeat := repeat.Repeat("z")
	fmt.Println(repeat)
	//output: zzzzz
}
