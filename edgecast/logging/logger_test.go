package logging

import (
	"testing"
)

func TestSample(t *testing.T) {

	got := 5
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
