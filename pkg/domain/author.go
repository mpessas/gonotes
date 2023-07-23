package domain

import "strings"

type Author string

func newAuthor(emailAddress string) *Author {
	a := Author(strings.ToLower(emailAddress))
	return &a
}
