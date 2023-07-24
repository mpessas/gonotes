package note

import (
	"resources/pkg/domain"
	"testing"

	"github.com/GoWebProd/uuid7"
)

type MemNoteRepository struct {
	Notes []Note
}

func (r *MemNoteRepository) Add(note *Note) {
	r.Notes = append(r.Notes, *note)
}

func TestCreateNotePortAuthorCase(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  []string
	}{
		{
			"Success",
			[]string{"This is a note", "john@example.org"},
			[]string{"This is a note", "john@example.org"},
		},
		{
			"Uppercase author",
			[]string{"This is a note", "John@example.org"},
			[]string{"This is a note", "john@example.org"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MemNoteRepository{}

			ans := CreateNotePort(&r, tt.input[0], tt.input[1], map[domain.TagKey]string{})
			if ans.Body != tt.want[0] || ans.Author != domain.Author(tt.want[1]) {
				t.Errorf("got %s, want %s", *ans, tt.want)
			}
		})
	}
}

func TestCreateNotePortNodeId(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
	}{
		{"Success", []string{"This is a note", "john@example.org"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MemNoteRepository{}

			ans := CreateNotePort(&r, tt.input[0], tt.input[1], map[domain.TagKey]string{})
			if _, err := uuid7.Parse(ans.Id.StringUUID()); err != nil {
				t.Errorf("invalid ID %s", ans.Id)
			}

			if len(r.Notes) != 1 {
				t.Errorf("note not added in repository")
			}

			if r.Notes[0].Id != ans.Id {
				t.Errorf("difference between repo and returned value")
			}
		})
	}
}

func TestCreateNotePortAddToRepository(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		tags  map[domain.TagKey]string
	}{
		{
			"Success",
			[]string{"This is a note", "john@example.org"},
			map[domain.TagKey]string{
				domain.ClientId: "123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MemNoteRepository{}

			ans := CreateNotePort(&r, tt.input[0], tt.input[1], tt.tags)
			if len(r.Notes) != 1 {
				t.Errorf("note not added in repository")
			}

			if r.Notes[0].Id != ans.Id {
				t.Errorf("difference between repo and returned value")
			}
		})
	}
}

func TestCreateNotePortValues(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		tags  map[domain.TagKey]string
	}{
		{
			"Success",
			[]string{"This is a note", "john@example.org"},
			map[domain.TagKey]string{
				domain.ClientId: "123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MemNoteRepository{}

			ans := CreateNotePort(&r, tt.input[0], tt.input[1], tt.tags)
			if ans.Author != domain.Author(tt.input[1]) {
				t.Errorf("note has wrong author")
			}
			if ans.Body != tt.input[0] {
				t.Errorf("note has wrong body")
			}
			if len(ans.Tags) != 1 {
				t.Errorf("note has wrong number of tags: %v", ans.Tags)
			}
			val, ok := ans.Tags[domain.ClientId]
			if !ok {
				t.Errorf("note does not have the tag key: %s", domain.ClientId.String())
			}
			if val != "123" {
				t.Errorf("note has wrong value for tag key ClientId: %s", val)
			}
		})
	}
}
