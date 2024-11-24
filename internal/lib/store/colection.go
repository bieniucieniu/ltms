package store

import "github.com/bieniucieniu/ltms/internal/lib/utils"

type colection map[string]*namespace

func NewColection(ns namespace, nsCodes ...string) (colection, error) {
	col := colection{}
	if err := col.Set(ns, nsCodes...); err != nil {
		return nil, err
	}
	return col, nil
}

func MustNewColection(ns namespace, nsCodes ...string) colection {
	return utils.Must(NewColection(ns, nsCodes...))
}

func (c colection) Set(ns namespace, nsCodes ...string) error {
	if ns == nil {
		ns = namespace{}
	}
	for _, code := range nsCodes {
		if !codeRegex.MatchString(code) {
			return ErrInvalidKey
		}
	}
	for _, code := range nsCodes {
		c[code] = &ns
	}
	return nil
}

func (c colection) MustSet(ns namespace, nsCodes ...string) colection {
	return utils.Must(c, c.Set(ns, nsCodes...))
}

func (c colection) Remove(code string) {
	var ptr *namespace = nil
	for key, value := range c {
		if key == code {
			ptr = value
			break
		}
	}
	for key, value := range c {
		if value == ptr {
			delete(c, key)
		}
	}
}
