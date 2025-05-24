package errs

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTTPError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"message": message})
}

func BadRequestError(c *gin.Context, message string) {
	HTTPError(
		c,
		http.StatusBadRequest,
		message,
	)
}

func InternalServerError(c *gin.Context, err error) {
	log.Println("Internal Server Error:", err)

	HTTPError(
		c,
		http.StatusInternalServerError,
		"Internal Server Error",
	)
}
