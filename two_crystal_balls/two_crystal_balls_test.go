package twocrystalballs

import "testing"

func TestTwoCrystalBalls(t *testing.T) {
	t.Run("found the break point", func(t *testing.T) {
		breaks := []bool{false, false, false, true, true, true, true}
		got := TwoCrystalBalls(breaks)
		want := 3

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, breaks)
		}
	})
}
