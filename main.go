package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html_tmpl/*")
	r.Static("/statics","./statics")
	r.GET("/", Index)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/aes", AesDecrypt)
	r.POST("/rsa", RsaDecrypt)
	r.Run("127.0.0.1:9800") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}