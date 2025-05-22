package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (d Dictionary) Add(term, definition string) error {
	if _, ok := d[term]; ok {
		return ErrWordExists
	}
	d[term] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if _, ok := d[word]; !ok {
		return ErrNotFound
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Search(search string) (string, error) {
	def, ok := d[search]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}
