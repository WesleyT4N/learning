package maps

type Dictionary map[string]string

type DictionaryError string

const (
	ErrNotFound         = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("word already exists in dict")
	ErrWordDoesNotExist = DictionaryError("word does not exist")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return res, nil
}

func (d Dictionary) Add(word, val string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, val string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = val
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
