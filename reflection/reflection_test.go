package reflection

import "testing"

func TestWalking(t *testing.T) {
  
  expected := "Chris"
  var got []string

  x := struct {
    Name string
  } {expected}

  walk(x, func (input string) {
    got = append(got, input) 
  })

  if len(got) != 1 {
    t.Errorf("wrong number of function calls, got %d wanted %d",len(got),1)
  }

  if got[0] != expected {
    t.Errorf("got %q wanted %q",got[0],expected)
  }
  
}
