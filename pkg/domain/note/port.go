package note

import "resources/pkg/domain"

func CreateNotePort(r NoteRepository, body, author_s string, tags map[domain.TagKey]string) *Note {
	author := domain.NewAuthor(author_s)
	note := NoteFrom(body, *author, tags)
	r.Add(note)
	return note
}
