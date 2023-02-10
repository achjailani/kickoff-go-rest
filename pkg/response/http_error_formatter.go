package response

import "github.com/gin-gonic/gin"

// ErrorOutput is struct to hold values need
type ErrorOutput struct {
	Gin      *gin.Context
	Code     int         `json:"code"`
	Language string      `json:"language"`
	Message  string      `json:"message"`
	Errors   interface{} `json:"errors"`
}

// ErrorResponse is struct type of base error response
type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

// NewHTTPError is a constructor of HTTP Success response
func NewHTTPError(c *gin.Context, errorMessage string, errors ...interface{}) *ErrorOutput {
	errorHTTPCode := c.Writer.Status()

	errOutput := &ErrorOutput{
		Gin:     c,
		Code:    errorHTTPCode,
		Message: errorMessage,
		Errors:  errors,
	}

	return errOutput
}

// SetEnLanguage is a function to set en language
func (e *ErrorOutput) SetEnLanguage() *ErrorOutput {
	e.Language = "en"

	return e
}

// JSON is a function to format json response
func (e *ErrorOutput) JSON() {
	e.SetEnLanguage()

	errorResponse := &ErrorResponse{
		Code:    e.Code,
		Message: e.Message,
		Errors:  e.Errors,
	}

	e.Gin.Header("Accept-Language", e.Language)
	e.Gin.JSON(e.Code, errorResponse)
}
