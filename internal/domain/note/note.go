package note

import "resources/internal/domain/tags"

type Note struct {
	Id     NoteId                 `json:"id"`
	Body   string                 `json:"body"`
	Author Author                 `json:"author"`
	Tags   map[tags.TagKey]string `json:"tags"`
}

func NoteFrom(body string, author Author, tags map[tags.TagKey]string) *Note {
	id := GenerateId()
	return &Note{id, body, author, tags}
}
