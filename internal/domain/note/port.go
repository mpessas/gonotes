package note

import (
	"resources/internal/domain/tags"
)

func CreateNotePort(r NoteRepository, body, author_s string, tags map[tags.TagKey]string) *Note {
	author := NewAuthor(author_s)
	note := NoteFrom(body, *author, tags)
	r.Add(note)
	return note
}

func SearchNotesPort(r NoteRepository, tags map[tags.TagKey]string) []Note {
	return r.Search(tags)
}
