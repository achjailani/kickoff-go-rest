package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/pkg/constant"
	"github.com/achjailani/kickoff-go-rest/pkg/exception"
	"github.com/achjailani/kickoff-go-rest/pkg/response"
	"github.com/achjailani/kickoff-go-rest/pkg/security"
	"net/http"
	"reflect"
	"regexp"
)

// UserPayload is
type UserPayload struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResponse is a struct
type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Validate is a function to validate request
func (u UserPayload) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&u.Username, validation.Required, validation.Length(7, 30), validation.Match(regexp.MustCompile("^[A-Za-z][A-Za-z0-9_]{7,30}$"))),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 100)),
	)
}

// UserCreateHandler is function to handle creation data
func (h *Handler) UserCreateHandler(c *gin.Context) {
	var payload UserPayload

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

	_, exists := h.repository.UserRepository.FindByUsername(c, payload.Username)
	if exists == nil {
		c.Status(http.StatusUnprocessableEntity)
		response.NewHTTPError(c, exception.ErrorUnprocessableEntity.Error(), "username has already been taken").JSON()
		return
	}

	role, _ := h.repository.RoleRepository.FindByCode(c, constant.RoleDefaultUser)

	user := entity.User{
		RoleID:   role.ID,
		Name:     payload.Name,
		Username: payload.Username,
		Password: security.HashMake(payload.Password),
	}
	err := h.repository.UserRepository.Save(c, &user)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		response.NewHTTPError(c, exception.ErrorUnprocessableEntity.Error()).JSON()
		return
	}

	response.NewHTTPSuccess(c, &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, fmt.Sprintf("OK")).JSON()
	return
}

// UserGetByIDHandler is function to handle single retrieve data
func (h *Handler) UserGetByIDHandler(c *gin.Context) {
	id := c.Param("id")

	if reflect.ValueOf(id).IsZero() {
		c.Status(http.StatusBadRequest)
		response.NewHTTPError(c, exception.ErrorBadRequest.Error(), []string{"missing id parameter"}).JSON()
		return
	}

	user, err := h.repository.UserRepository.FindByID(c, id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		response.NewHTTPError(c, err.Error(), []string{}).JSON()
		return
	}

	response.NewHTTPSuccess(c, &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, "Ok").JSON()
	return
}

// UserListHandler is function to handle list of data
func (h *Handler) UserListHandler(c *gin.Context) {
	var users []*UserResponse

	results, err := h.repository.UserRepository.FindAll(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		response.NewHTTPError(c, exception.ErrorInternalServerError.Error(), err.Error()).JSON()
		return
	}

	for _, row := range results {
		users = append(users, &UserResponse{
			ID:        row.ID,
			Name:      row.Name,
			Username:  row.Username,
			CreatedAt: row.CreatedAt.String(),
			UpdatedAt: row.UpdatedAt.String(),
		})
	}

	response.NewHTTPSuccess(c, users, fmt.Sprintf("OK")).JSON()
	return
}
