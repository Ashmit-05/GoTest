package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
  buffer := &bytes.Buffer{}

  Countdown(buffer)
  
}
