package maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Get(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", d.KeyError(key)
	}
	return definition, nil
}

func (d Dictionary) Set(key, value string) {
	d[key] = value
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Get(key)
	switch err {
	case nil:
		return d.ExistsError(key)
	case err:
		keyErr := d.KeyError(key)
		isKeyErr := (keyErr.Error() == err.Error())
		if isKeyErr {
			d[key] = value
		}
	}
	return nil
}

func (d Dictionary) KeyError(key string) error {
	errorMsg := fmt.Sprintf("%s not in dictionary", key)
	return errors.New(errorMsg)
}

func (d Dictionary) ExistsError(key string) error {
	errorMsg := fmt.Sprintf("%s already exists in dictionary", key)
	return errors.New(errorMsg)
}
