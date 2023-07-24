package note

import "resources/pkg/domain"

func CreateNotePort(r NoteRepository, body, author_s string) *Note {
	author := domain.NewAuthor(author_s)
	note := NoteFrom(body, *author)
	r.Add(note)
	return note
}
