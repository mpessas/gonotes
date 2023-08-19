package main

import (
	"github.com/gin-gonic/gin"
	"resources/internal/http"
)

func main() {
	r := gin.Default()
	http.SetupRoutes(r)
	r.Run()
}
