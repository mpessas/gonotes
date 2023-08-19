package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resources/internal/domain/note"
	"resources/internal/domain/tags"
)

type MemNoteRepository struct {
	Notes []note.Note
}

func (r *MemNoteRepository) Add(note *note.Note) {
	r.Notes = append(r.Notes, *note)
}

func (r *MemNoteRepository) Search(tags map[tags.TagKey]string) []note.Note {
	var v string
	var ok bool
	var isMatch bool
	var matches []note.Note

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

type CreateNotePayload struct {
	Body   string                 `json:"body" binding:"required"`
	Author string                 `json:"author" binding:"required,email"`
	Tags   map[tags.TagKey]string `json:"tags"`
}

func createNote(c *gin.Context) {
	var payload CreateNotePayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	repo := MemNoteRepository{}
	fmt.Printf("%v", payload.Tags)
	rv := note.CreateNotePort(&repo, payload.Body, payload.Author, payload.Tags)
	c.JSON(http.StatusCreated, rv)
}
