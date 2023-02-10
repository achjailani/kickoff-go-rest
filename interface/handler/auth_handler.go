package handler

import (
	"fmt"
	"github.com/achjailani/kickoff-go-rest/pkg/exception"
	"github.com/achjailani/kickoff-go-rest/pkg/response"
	"github.com/achjailani/kickoff-go-rest/pkg/security"
	"github.com/achjailani/kickoff-go-rest/pkg/security/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// AuthPayload is a struct defines auth payload
type AuthPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate is function to validate AuthPayload
func (a AuthPayload) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required, validation.Length(6, 100)),
		validation.Field(&a.Password, validation.Required, validation.Length(8, 70)),
	)
}

// AuthResponse is a struct defines auth response
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiredIn   int64  `json:"expired_in"`
}

// Login is a function to handle authentication
func (h *Handler) Login(c *gin.Context) {
	var payload AuthPayload

	if err := c.ShouldBind(&payload); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		response.NewHTTPError(c, exception.ErrorUnprocessableEntity.Error()).JSON()
		return
	}

	errV := payload.Validate()
	if errV != nil {
		c.Status(http.StatusUnprocessableEntity)
		response.NewHTTPError(c, exception.ErrorUnprocessableEntity.Error(), errV).JSON()
		return
	}

	user, errFind := h.repository.UserRepository.FindByUsername(c, payload.Username)
	if errFind != nil {
		c.Status(http.StatusUnauthorized)
		response.NewHTTPError(c, fmt.Sprintf("Invalid username or password")).JSON()
		return
	}

	if ok := security.HashVerify(payload.Password, user.Password); !ok {
		c.Status(http.StatusUnauthorized)
		response.NewHTTPError(c, fmt.Sprintf("Invalid username or password")).JSON()
		return
	}

	// define ttl
	ttl := time.Hour * 24

	// define jwt payload
	jwtPayload := jwt.JWTPayload{
		Identifier: user.ID,
		Name:       user.Name,
	}

	token, err := h.jwt.Create(ttl, &jwtPayload)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		response.NewHTTPError(c, fmt.Sprintf("something went wrong.")).JSON()
		return
	}

	expiredIn := time.Now().Add(ttl).Unix()

	result := AuthResponse{
		AccessToken: token,
		ExpiredIn:   expiredIn,
	}

	response.NewHTTPSuccess(c, result, fmt.Sprintf("Ok")).JSON()
}
