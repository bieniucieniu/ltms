package store

import "github.com/bieniucieniu/ltms/internal/lib/utils"

type Store struct {
	colections map[string]*colection
}

func New() *Store {
	return &Store{colections: map[string]*colection{}}
}

func (s *Store) Set(col colection, codes ...string) error {
	if col == nil {
		col = colection{}
	}
	for _, code := range codes {
		if !codeRegex.MatchString(code) {
			return ErrInvalidKey
		}
	}
	for _, code := range codes {
		s.colections[code] = &col
	}
	return nil
}

func (s *Store) MustSet(col colection, codes ...string) *Store {
	return utils.Must(s, s.Set(col, codes...))
}

func (s *Store) Remove(code string) {
	var ptr *colection = nil
	for key, value := range s.colections {
		if key == code {
			ptr = value
			break
		}
	}
	for key, value := range s.colections {
		if value == ptr {
			delete(s.colections, key)
		}
	}
}

func (s *Store) SetNamespace(ns namespace, colection string, nsCodes ...string) error {
	c := s.colections[colection]
	if c == nil {
		return ErrNoSuchColection
	}
	c.Set(ns, nsCodes...)
	return nil
}

func (s *Store) MustSetNamespace(ns namespace, colection string, nsCodes ...string) *Store {
	return utils.Must(s, s.SetNamespace(ns, colection, nsCodes...))
}
