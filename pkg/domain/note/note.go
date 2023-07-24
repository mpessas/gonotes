package note

import "resources/pkg/domain"

type Note struct {
	Id     NoteId        `json:"id"`
	Body   string        `json:"body"`
	Author domain.Author `json:"author"`
}

func NoteFrom(body string, author domain.Author) *Note {
	id := GenerateId()
	return &Note{id, body, author}
}
