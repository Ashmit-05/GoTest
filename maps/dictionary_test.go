package maps

import "testing"

func TestSearch(t *testing.T) {

  t.Run("known word",func(t *testing.T) {
    dictionary := Dictionary{"test" : "this is just a test"}

    got,_ := dictionary.Search("test")
    want := "this is just a test"

    assertString(t, got, want)
  })

  t.Run("unknown word",func(t *testing.T) {
    dictionary := Dictionary{"test" : "this is just a test"}

    _,err := dictionary.Search("something")
    want := NotFoundError

    if err == nil {
      t.Fatal("expected an error but didn't get one")
    }

    assertError(t,err, want)
  })
}

func assertString(t testing.TB, got, want string) {
  t.Helper()

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

func assertError(t testing.TB, got, want error) {
  t.Helper()

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
