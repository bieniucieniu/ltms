package store

import (
	"regexp"
)

var codeRegex = regexp.MustCompile(`^[\w_]+$`)
var keyRegex = regexp.MustCompile(`^[a-z0-9]{1,}(?:\.[a-z0-9]{1,})*$`)
