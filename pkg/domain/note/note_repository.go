package note

type NoteRepository interface {
	Add(note *Note)
}
