package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hola Mundo."
	
	if got := Hello(); want != got {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}