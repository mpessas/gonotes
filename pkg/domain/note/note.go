package note

import "resources/pkg/domain"

type Note struct {
	Id     NoteId                   `json:"id"`
	Body   string                   `json:"body"`
	Author domain.Author            `json:"author"`
	Tags   map[domain.TagKey]string `json:"tags"`
}

func NoteFrom(body string, author domain.Author, tags map[domain.TagKey]string) *Note {
	id := GenerateId()
	return &Note{id, body, author, tags}
}
