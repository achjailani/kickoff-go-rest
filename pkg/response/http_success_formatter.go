package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessOutput is struct to hold values need
type SuccessOutput struct {
	Gin      *gin.Context
	Data     interface{} `json:"data"`
	Language string      `json:"language"`
	Message  string      `json:"message"`
}

// SuccessResponse is struct type of base error response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewHTTPSuccess is a constructor of HTTP Success response
func NewHTTPSuccess(c *gin.Context, data interface{}, message string) *SuccessOutput {
	return &SuccessOutput{
		Gin:     c,
		Data:    data,
		Message: message,
	}
}

// SetEnLanguage is a function to set en language
func (e *SuccessOutput) SetEnLanguage() {
	e.Language = "en"
}

// JSON is a function to format json response
func (e *SuccessOutput) JSON() {
	e.SetEnLanguage()

	response := &SuccessResponse{
		Message: e.Message,
		Data:    e.Data,
	}

	e.Gin.Header("Accept-Language", e.Language)
	e.Gin.JSON(http.StatusOK, response)
}
