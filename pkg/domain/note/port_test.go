package note

import (
	"github.com/GoWebProd/uuid7"
	"resources/pkg/domain"
	"testing"
)

type MemNoteRepository struct {
	Notes []Note
}

func (r *MemNoteRepository) Add(note *Note) {
	r.Notes = append(r.Notes, *note)
}

func TestCreateNotePort(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  []string
	}{
		{"Success", []string{"This is a note", "john@example.org"}, []string{"This is a note", "john@example.org"}},
		{"Uppercase author", []string{"This is a note", "John@example.org"}, []string{"This is a note", "john@example.org"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MemNoteRepository{}

			ans := CreateNotePort(&r, tt.input[0], tt.input[1])
			if ans.Body != tt.want[0] || ans.Author != domain.Author(tt.want[1]) {
				t.Errorf("got %s, want %s", *ans, tt.want)
			}
			if _, err := uuid7.Parse(ans.Id.StringUUID()); err != nil {
				t.Errorf("invalid ID %s", ans.Id)
			}

			if len(r.Notes) != 1 {
				t.Errorf("note not added in repository")
			}

			if r.Notes[0] != *ans {
				t.Errorf("difference between repo and returned value")
			}
		})
	}
}
