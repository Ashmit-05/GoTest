package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer,"Tyler")

  got := buffer.String()
  want := "Hello, Tyler"

  if got != want {
    t.Errorf("got %q want %q",got, want)
  }
  
}
