package domain

import "github.com/GoWebProd/uuid7"

type Note struct {
	Id     string `json:"id"`
	Body   string `json:"body"`
	Author Author `json:"author"`
}

func newNote(body string, author Author) *Note {
	id := uuid7.New().Next().String()
	return &Note{id, body, author}
}

func CreateNotePort(body, author_s string) *Note {
	author := newAuthor(author_s)
	return newNote(body, *author)
}
