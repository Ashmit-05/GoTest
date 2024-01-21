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

func TestAdd(t *testing.T)  {
  t.Run("new word", func(t *testing.T) {
    dictionary := Dictionary{}
    err := dictionary.Add("test","this is just a test") 

    assertError(t, err,nil)
    assertAddition(t, dictionary,"test","this is just a test")
  })

  t.Run("existing word", func(t *testing.T) {
    dictionary := Dictionary{"test":"this is just a test"}
    err := dictionary.Add("test","this is just a test")

    assertError(t,err, ErrWordExists)
    assertAddition(t, dictionary, "test","this is just a test")
  })
}

func TestUpdate(t *testing.T)  {
  t.Run("new word",func(t *testing.T) {
    dictionary := Dictionary{"test":"this is just a test"}
    err := dictionary.Update("some","thing")

    assertError(t,err, ErrWordDoesNotExist)
  })
   
  t.Run("existing word", func(t *testing.T) {
    dictionary := Dictionary{"test":"this is just a test"}
    err := dictionary.Update("test","this is updated")

    assertError(t,err, nil)
    assertAddition(t, dictionary, "test","this is updated")
  })
}

func TestDelete(t *testing.T) {
  t.Run("word exists", func(t *testing.T) {
    dictionary := Dictionary{"test":"this is just a test"}
    err := dictionary.Delete("test")

    assertError(t, err, nil)
    _,err = dictionary.Search("test")
    assertError(t,err , NotFoundError)
  })
  
  t.Run("word does not exist", func(t *testing.T) {
    dictionary := Dictionary{"test":"this is just a test"}
    err := dictionary.Delete("rest")

    assertError(t, err, NotFoundError)
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

func assertAddition(t testing.TB,d Dictionary, key, value string) {
  t.Helper()

  got, err := d.Search(key)
  if err != nil {
    t.Fatal("should find added word :", err)
  }

  assertString(t, got, value)
}
