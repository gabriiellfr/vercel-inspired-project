// app/handlers/index_handler.go

package handlers

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.File("./templates/index.html")
}
