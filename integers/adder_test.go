package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(3, 4)
	want := 7

	if got != want {
		t.Errorf("got %q expected %q", got, want)
	}

}
