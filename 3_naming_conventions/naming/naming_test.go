package naming

import "testing"

func TestColo(t *testing.T) {
	arg := "blue"
	want := "#0000FF"
	got := Color(arg)

	if got != want {
		t.Errorf("Color(%q) = %q; want = %q", arg, got, want)
	}
}

func TestDog(t *testing.T) {

}

func TestSpeak(t *testing.T) {

}
