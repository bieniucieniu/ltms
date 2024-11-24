package store

import "github.com/bieniucieniu/ltms/internal/lib/utils"

type namespace map[string]string

func NewNamespace(init map[string]string) (namespace, error) {
	ns := namespace{}
	for key, value := range init {
		if !keyRegex.MatchString(key) {
			return nil, ErrInvalidKey
		}
		ns[key] = value
	}

	return ns, nil
}

func MustNewNamespace(init map[string]string) namespace {
	return utils.Must(NewNamespace(init))
}

func (ns namespace) Set(key, value string) error {
	if !keyRegex.MatchString(key) {
		return ErrInvalidKey
	}
	ns[key] = value
	return nil
}

func (ns namespace) MustSet(key, value string) namespace {
	if err := ns.Set(key, value); err != nil {
		panic(err)
	}
	return ns
}

func (ns namespace) Remove(keys ...string) {
	for _, key := range keys {
		delete(ns, key)
	}
}
