package maps

type Dictionary map[string]string

type DictionaryErr string

const (
  NotFoundError = DictionaryErr("could not find the word you are looking for")
  ErrWordExists = DictionaryErr("word already exists in the dictionary")
  ErrWordDoesNotExist = DictionaryErr("could not update because word already exists")
)

func (e DictionaryErr) Error() string {
  return string(e)
}

func (d Dictionary) Search(word string) (string,error) {
  if d[word] == "" {
    return "", NotFoundError
  }
  return d[word], nil
}

func (d Dictionary) Add(key, value string) error {
  _, err := d.Search(key)

  switch err {
  case NotFoundError:
    d[key] = value
  case nil :
    return ErrWordExists
  default :
    return err
  }

  return nil
}

func (d Dictionary) Update(key, value string) error {
  _, err := d.Search(key)

  switch err {
  case NotFoundError:
    return ErrWordDoesNotExist
  case nil :
    d[key] = value
  default :
    return err
  }

  return nil
}

func (d Dictionary) Delete(key string) error {
  _, err := d.Search(key)

  switch err {
  case NotFoundError:
    return err
  case nil :
    delete(d,key)
  default :
    return err
  }

  return nil
}
