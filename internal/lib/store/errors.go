package store

import "errors"

var ErrInvalidKey = errors.New("invalid key")
var ErrNoSuchColection = errors.New("no such colection")
var ErrNoSuchNamespace = errors.New("no such namespace")
