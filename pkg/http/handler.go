package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resources/pkg/domain"
)

type CreateNotePayload struct {
	Body   string `json:"body" binding:"required"`
	Author string `json:"author" binding:"required,email"`
}

func createNote(c *gin.Context) {
	var payload CreateNotePayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	rv := domain.CreateNotePort(payload.Body, payload.Author)
	c.JSON(http.StatusCreated, rv)
}
