package note

import "strings"

type Author string

func NewAuthor(emailAddress string) *Author {
	a := Author(strings.ToLower(emailAddress))
	return &a
}
