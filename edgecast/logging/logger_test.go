package logging

import (
	"testing"
)

func TestSample(t *testing.T) {

	got := 10
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
