package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/pkg/exception"
	"github.com/achjailani/kickoff-go-rest/pkg/response"
	"net/http"
)

type RoleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type CreateRoleRequest struct {
	Name        string
	Code        string
	Description string
}

func (h *Handler) CreateRoleHandler(c *gin.Context) {
	var role entity.Role

	if err := c.ShouldBind(&role); err != nil {
		c.Status(http.StatusBadRequest)
		response.NewHTTPError(c, exception.ErrorBadRequest.Error()).JSON()
		return
	}

	err := h.repository.RoleRepository.Save(c, &role)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		response.NewHTTPError(c, fmt.Sprintf("%v", err)).JSON()
		return
	}

	response.NewHTTPSuccess(c, role, fmt.Sprintf("Ok")).JSON()
	return
}

func (h *Handler) GetRoleHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.Status(http.StatusBadRequest)
		response.NewHTTPError(c, fmt.Sprintf("Invalid parameter."))
		return
	}

	role, err := h.repository.RoleRepository.FindByID(c, id)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		response.NewHTTPError(c, exception.ErrorUnprocessableEntity.Error())
		return
	}

	response.NewHTTPSuccess(c, role, fmt.Sprintf("OK")).JSON()
	return
}

func (h *Handler) GetAllRoleHandler(c *gin.Context) {
	roles, err := h.repository.RoleRepository.FindAll(c)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		response.NewHTTPError(c, exception.ErrorUnprocessableEntity.Error())
		return
	}

	response.NewHTTPSuccess(c, roles, fmt.Sprintf("OK")).JSON()
	return
}
