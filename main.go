package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	//os.Setenv("GIN_MODE", "release")
	//gin.SetMode(gin.ReleaseMode)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.StaticFS("/templates", http.Dir(filepath.Join(os.Getenv("GOPATH"),
		"src/github.com/ameykpatil/gospike/templates")))
	router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"),
		"src/github.com/ameykpatil/gospike/templates/*"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "GoSpike",
		})
	})
	router.GET("/connect", Connect)
	router.GET("/record/:key", GetRecord)
	router.POST("/record", AddRecord)
	//router.PUT("/record/:key", updateRecord)
	router.DELETE("/record/:key", DeleteRecord)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":4848")
}
