package maps

import "errors"

type Dictionary map[string]string

var NotFound = errors.New("No such key")

func (d Dictionary) Search(key string) (string, error) {
	def, ok := d[key]

	if !ok {
		return def, NotFound
	}

	return def, nil

}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
