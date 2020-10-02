package utils

import (
	"strings"
)

func IsValidType(e string) bool {
	e = strings.ToLower(e)
	return e == "company" || e == "person" || e == "article" 
}
