package main

const (
	ErrWordNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("could not add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("could not perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]
	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)
	switch err {
	case ErrWordNotFound:
		d[term] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(term, definition string) error {
	_, err := d.Search(term)
	switch err {
	case nil:
		d[term] = definition
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(term string) error {
	_, err := d.Search(term)
	switch err {
	case nil:
		delete(d, term)
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}
