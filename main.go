package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type URLObject struct {
	contenttype string
	content string
}

func main() {
	URLPool := make(map[string]URLObject)
	server := gin.Default()
	server.LoadHTMLGlob("view/*")
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	server.GET("/:parameter", func(c *gin.Context) {
		parameter := c.Param("parameter")
		_, ParameterExist := URLPool[parameter]
		if ParameterExist {
			contenttype := URLPool[parameter].contenttype
			content := URLPool[parameter].content
			if  contenttype != "text" {
				c.String(http.StatusOK, fmt.Sprintf(content))
			} else if contenttype != "url"{
				c.Redirect(http.StatusMovedPermanently, content)
			}
		} else {
			c.String(http.StatusOK, "No paramter")
		}
	})

	server.POST("/", func(c *gin.Context) {
		parameter := c.PostForm("content_parameter")
		content := c.PostForm("content")
		contenttype := c.PostForm("content_type")
		fmt.Printf("%s, %s, %s", parameter, content, contenttype)
		URLPool[parameter] = URLObject{contenttype, content}
		c.String(http.StatusOK, "ok")
	})
	server.Run()
}
