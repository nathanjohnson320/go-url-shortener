package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Url struct {
	Id       int    `json:"id"`
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

func main() {
	r := gin.Default()

	r.Static("/assets", "./dist/assets")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r.GET("/urls", func(c *gin.Context) {
		urls := []Url{}
		c.JSON(http.StatusOK, urls)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
