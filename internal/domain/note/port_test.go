package note

import (
	"github.com/GoWebProd/uuid7"
	"resources/internal/domain/tags"
	"testing"
)

type MemNoteRepository struct {
	Notes []Note
}

func (r *MemNoteRepository) Add(note *Note) {
	r.Notes = append(r.Notes, *note)
}

func (r *MemNoteRepository) Search(tags map[tags.TagKey]string) []Note {
	var v string
	var ok bool
	var isMatch bool
	var matches []Note

	for _, n := range r.Notes {
		isMatch = true
		for tagKey, tagValue := range tags {
			v, ok = n.Tags[tagKey]
			if !ok {
				isMatch = false
				break
			}
			if tagValue != v {
				isMatch = false
				break
			}
		}
		if isMatch {
			matches = append(matches, n)
		}
	}
	return matches
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

			ans := CreateNotePort(&r, tt.input[0], tt.input[1], map[tags.TagKey]string{})
			if ans.Body != tt.want[0] || ans.Author != Author(tt.want[1]) {
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

			ans := CreateNotePort(&r, tt.input[0], tt.input[1], map[tags.TagKey]string{})
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
		tags  map[tags.TagKey]string
	}{
		{
			"Success",
			[]string{"This is a note", "john@example.org"},
			map[tags.TagKey]string{
				tags.ClientId: "123",
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
		tags  map[tags.TagKey]string
	}{
		{
			"Success",
			[]string{"This is a note", "john@example.org"},
			map[tags.TagKey]string{
				tags.ClientId: "123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MemNoteRepository{}

			ans := CreateNotePort(&r, tt.input[0], tt.input[1], tt.tags)
			if ans.Author != Author(tt.input[1]) {
				t.Errorf("note has wrong author")
			}
			if ans.Body != tt.input[0] {
				t.Errorf("note has wrong body")
			}
			if len(ans.Tags) != 1 {
				t.Errorf("note has wrong number of tags: %v", ans.Tags)
			}
			val, ok := ans.Tags[tags.ClientId]
			if !ok {
				t.Errorf("note does not have the tag key: %s", tags.ClientId.String())
			}
			if val != "123" {
				t.Errorf("note has wrong value for tag key ClientId: %s", val)
			}
		})
	}
}

func TestSearchPort(t *testing.T) {
	note1 := Note{
		GenerateId(),
		"This is a note",
		Author("john@example.org"),
		map[tags.TagKey]string{
			tags.ClientId: "123",
			tags.Carrier:  "Progressive",
		},
	}
	note2 := Note{
		GenerateId(),
		"This is a second note",
		Author("john@example.org"),
		map[tags.TagKey]string{
			tags.ClientId: "124",
			tags.Carrier:  "Progressive",
		},
	}
	note3 := Note{
		GenerateId(),
		"This is a third note",
		Author("jack@example.org"),
		map[tags.TagKey]string{
			tags.ClientId: "123",
			tags.Carrier:  "Progressive",
		},
	}
	note4 := Note{
		GenerateId(),
		"This is a fourth note",
		Author("jane@example.org"),
		map[tags.TagKey]string{
			tags.ClientId: "123",
			tags.Carrier:  "Humana",
		},
	}

	r := MemNoteRepository{}
	r.Add(&note1)
	r.Add(&note2)
	r.Add(&note3)
	r.Add(&note4)

	var tests = []struct {
		name  string
		input map[tags.TagKey]string
		want  []Note
	}{
		{
			"Success",
			map[tags.TagKey]string{
				tags.ClientId: "123",
				tags.Carrier:  "Progressive",
			},
			[]Note{note1, note3},
		},
		{
			"No matches because of ClientId",
			map[tags.TagKey]string{
				tags.ClientId: "125",
				tags.Carrier:  "Progressive",
			},
			[]Note{},
		},
		{
			"No matches because of Carrier",
			map[tags.TagKey]string{
				tags.ClientId: "124",
				tags.Carrier:  "Humana",
			},
			[]Note{},
		},
		{
			"No matches because of both",
			map[tags.TagKey]string{
				tags.ClientId: "125",
				tags.Carrier:  "Lumico",
			},
			[]Note{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := SearchNotesPort(&r, tt.input)
			if len(ans) != len(tt.want) {
				t.Errorf("non-matching lengths, got %v, wanted %v", ans, tt.want)
			}
			// if ans != tt.want {
			// 	t.Errorf("got %v, wanted %v", ans, want)
			// }
		})
	}
}
