package domain

import (
	"github.com/GoWebProd/uuid7"
	"testing"
)

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
			ans := CreateNotePort(tt.input[0], tt.input[1])
			if ans.Body != tt.want[0] || ans.Author != Author(tt.want[1]) {
				t.Errorf("got %s, want %s", *ans, tt.want)
			}
			if _, err := uuid7.Parse(ans.Id); err != nil {
				t.Errorf("invalid ID %s", ans.Id)
			}

		})
	}
}
