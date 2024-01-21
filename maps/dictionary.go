package maps

import "errors"

type Dictionary map[string]string

var NotFoundError = errors.New("could not find the word you are looking for")

func (d Dictionary) Search(word string) (string,error) {
  if d[word] == "" {
    return "", NotFoundError
  }
  return d[word], nil
}
