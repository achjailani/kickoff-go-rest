package handler

import (
	"fmt"
	"github.com/achjailani/kickoff-go-rest/infrastructure/core/provider/connection"
	"github.com/achjailani/kickoff-go-rest/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	DB string `json:"db"`
}

func (e *Handler) Ping(c *gin.Context) {
	var pingResponse PingResponse

	_, errDBConn := connection.NewDBConnection(e.config)
	if errDBConn != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errDBConn)

		pingResponse.DB = "No Ok"
	} else {
		pingResponse.DB = "Ok"
	}

	response.NewHTTPSuccess(c, pingResponse, fmt.Sprintf("OK")).JSON()
	return
}
