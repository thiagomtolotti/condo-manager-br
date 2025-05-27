package errs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (err *AppError) Error() string {
	return fmt.Sprintf("%s: %v", err.Message, err.Err)
}

func HandleError(c *gin.Context, err *AppError) {
	if err.Code >= http.StatusInternalServerError {
		if err != nil {
			log.Printf("Internal Error: %v\n", err.Error())
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
		Err:     err,
	}
}

func BadRequest(message string, err error) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
		Err:     err,
	}
}
