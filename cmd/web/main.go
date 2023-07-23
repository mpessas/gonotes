package main

import (
	"github.com/gin-gonic/gin"
	"resources/pkg/http"
)

func main() {
	r := gin.Default()
	http.SetupRoutes(r)
	r.Run()
}
