package errs

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int
	Message string
	Error   error
}

func HandleError(c *gin.Context, err *AppError) {
	if err.Code >= http.StatusInternalServerError {
		if err.Error != nil {
			log.Printf("Internal Error: %v\n", err.Error)
		} else {
			log.Println("No error provided")
		}
	}

	c.JSON(err.Code, gin.H{
		"message": err.Message,
	})
}

func Unexpected(err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   err,
	}
}

func BadRequest(message string, err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Error:   err,
	}
}
