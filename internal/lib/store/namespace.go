package store

import (
	"errors"
	"regexp"
)

var translationKeyRegex = regexp.MustCompile(`^[a-z0-9]{1,}(?:\.[a-z0-9]{1,})*$`)

type Namespace struct {
	records map[string]string
}

func NewNamespace(init map[string]string) (*Namespace, error) {
	ns := Namespace{
		records: map[string]string{},
	}
	if init != nil {
		for key := range init {
			if !translationKeyRegex.MatchString(key) {
				return nil, errors.Join(ErrInvalidKey, errors.New(key))
			}
		}
		ns.records = init
	}

	return &ns, nil
}

func (t *Namespace) AddRecord(key, value string) error {
	if !translationKeyRegex.MatchString(key) {
		return ErrInvalidKey
	}
	t.records[key] = value
	return nil
}
func (t *Namespace) AddRecords(values map[string]string) error {
	for key, value := range values {
		if !translationKeyRegex.MatchString(key) {
			return errors.Join(ErrInvalidKey, errors.New(key))
		}
		t.records[key] = value
	}
	return nil
}

func (t *Namespace) RemoveRecord(key string) {
	delete(t.records, key)
}

func (t *Namespace) RemoveRecords(keys []string) {
	for _, key := range keys {
		delete(t.records, key)
	}
}
