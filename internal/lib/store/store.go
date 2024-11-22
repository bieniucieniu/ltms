package store

import (
	"regexp"
)

type Srore struct {
	colections map[string]*Colection
}

func NewStore() *Srore {
	return &Srore{colections: map[string]*Colection{}}
}

var codeRegex = regexp.MustCompile(`^[\w_]+$`)

func (c *Srore) AddColection(codes []string, colection *Colection) error {
	for _, code := range codes {
		if !codeRegex.MatchString(code) {
			return ErrInvalidKey
		}
	}
	for _, code := range codes {

		c.colections[code] = colection
	}
	return nil
}

func (c *Srore) RemoveColection(code string, colection *Colection) error {
	var ptr *Colection
	for key, value := range c.colections {
		if key == code {
			ptr = value
			break
		}
	}
	for key, value := range c.colections {
		if value == ptr {
			delete(c.colections, key)
		}
	}

	return nil
}
