package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/achjailani/kickoff-go-rest/pkg/response"
	"log"
	"net/http"
	"time"
)

type DevResponse struct {
	Message string `json:"message"`
}

func (h *Handler) DevErrorHandler(c *gin.Context) {
	devResponse := DevResponse{
		Message: fmt.Sprintf("Hi, this error is testing, time: %v", time.Now()),
	}

	c.Status(http.StatusBadRequest)
	response.NewHTTPError(c, devResponse.Message).JSON()
	return
}

func (h *Handler) DevSuccessHandler(c *gin.Context) {
	devResponse := DevResponse{
		Message: fmt.Sprintf("Hi, this message is testing, time: %v", time.Now()),
	}

	response.NewHTTPSuccess(c, nil, devResponse.Message).JSON()
	return
}

// DevAsyncHandler is a function to run goroutine inside
func (h *Handler) DevAsyncHandler(c *gin.Context) {
	devResponse := DevResponse{
		Message: fmt.Sprintf("Hi, this one is async."),
	}

	// Use copy context for async process (goroutine)
	cCp := c.Copy()

	// Look at the console to see the results
	go func() {
		time.Sleep(time.Second * 5)
		log.Println("Done Async process! in path " + cCp.Request.URL.Path)
	}()

	response.NewHTTPSuccess(c, nil, devResponse.Message).JSON()
	return
}

// DevSyncHandler is a function to wait until 5 seconds
func (h *Handler) DevSyncHandler(c *gin.Context) {
	devResponse := DevResponse{
		Message: fmt.Sprintf("Hi, this one is sync."),
	}

	time.Sleep(5 * time.Second)

	log.Println("Done Sync process! in path " + c.Request.URL.Path)

	response.NewHTTPSuccess(c, nil, devResponse.Message).JSON()
	return
}
