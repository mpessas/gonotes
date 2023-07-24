package note

import "resources/pkg/domain/libs"

type NoteId string

const prefix = "note"

func GenerateId() NoteId {
	return NoteId(libs.GenerateId(prefix))
}

func (v *NoteId) String() string {
	return string(*v)
}

func (v *NoteId) StringUUID() string {
	return string(*v)[5:]
}
