package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eddywm/go-shortner/handler"
	"github.com/eddywm/go-shortner/store"
	"github.com/gin-gonic/gin"
)

func main() {

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"create": "POST create-short-url",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
