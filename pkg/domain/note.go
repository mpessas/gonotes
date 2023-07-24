package domain

import "github.com/GoWebProd/uuid7"

type NoteId string

func NoteIdFromUuid(v uuid7.UUID) NoteId {
	return NoteId("note_" + v.String())
}

func (v *NoteId) String() string {
	return string(*v)
}

func (v *NoteId) StringUUID() string {
	return string(*v)[5:]
}

type Note struct {
	Id     NoteId `json:"id"`
	Body   string `json:"body"`
	Author Author `json:"author"`
}

type NoteRepository interface {
	Add(note *Note)
}

func newNote(body string, author Author) *Note {
	id := NoteIdFromUuid(uuid7.New().Next())
	return &Note{id, body, author}
}

func CreateNotePort(r NoteRepository, body, author_s string) *Note {
	author := newAuthor(author_s)
	note := newNote(body, *author)
	r.Add(note)
	return note
}
