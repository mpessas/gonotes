package note

import "resources/internal/domain/tags"

type NoteRepository interface {
	Add(note *Note)
	Search(tags map[tags.TagKey]string) []Note
}
