package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Url struct {
	Id       int    `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func listUrls(db *sql.DB) ([]Url, error) {
	// Query the database
	rows, err := db.Query("SELECT * FROM urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	urls := []Url{}
	for rows.Next() {
		var url Url

		if err := rows.Scan(&url.Id, &url.LongUrl, &url.ShortUrl); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func main() {
	connStr := "user=postgres dbname=url_shortener sslmode=disable password=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Static("/assets", "./dist/assets")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r.GET("/urls", func(c *gin.Context) {
		urls, error := listUrls(db)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}

		c.JSON(http.StatusOK, urls)
	})

	r.POST("/urls", func(c *gin.Context) {
		var url Url
		if err := c.ShouldBindJSON(&url); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert the url into the database
		id, err := gonanoid.New()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rows, err := db.Query("INSERT INTO urls (long_url, short_url) VALUES ($1, $2) RETURNING *", url.LongUrl, id)
		rows.Next()
		rows.Scan(&url.Id, &url.LongUrl, &url.ShortUrl)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, url)
	})

	r.DELETE("/urls/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM urls WHERE short_url = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
