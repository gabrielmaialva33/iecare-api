package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
	"iecare-api/src/app/services"
	"iecare-api/src/app/validators"
	"strconv"
	"strings"
)

// RolesController is a controller for roles
type RolesController struct {
	rr services.IRoleServices
}

// NewRolesController is a constructor for RolesController
func NewRolesController(rr interfaces.RoleInterface) *RolesController {
	return &RolesController{rr}
}

func (r *RolesController) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")

	roles, err := r.rr.List(paginate.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Search:      search,
		Sort:        sort,
		Order:       order,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting roles",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(roles)
}

func (r *RolesController) Get(c *fiber.Ctx) error {
	uuid := c.Params("roleId")

	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	role, err := r.rr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	return c.JSON(role.PublicRole())
}

func (r *RolesController) Store(c *fiber.Ctx) error {
	data := models.Role{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing body",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	role := models.Role{
		Name: strings.ToLower(data.Slug),
	}
	if err := mergo.Merge(&role, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": true,
		})
	}

	if errors := validators.ValidateStruct(role); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	newRole, err := r.rr.Store(&role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while creating role",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	return c.JSON(newRole.PublicRole())
}

func (r *RolesController) Edit(c *fiber.Ctx) error {
	uuid := c.Params("roleId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	data := models.Role{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	role, err := r.rr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	dstRole := models.Role{
		Id:   role.Id,
		Name: strings.ToLower(data.Slug),
	}
	if err := mergo.Merge(&dstRole, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidatePartialStruct(dstRole); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	editedRole, err := r.rr.Edit(&dstRole)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while updating user",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(editedRole.PublicRole())
}

func (r *RolesController) Delete(c *fiber.Ctx) error {
	uuid := c.Params("roleId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	role, err := r.rr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	deleteRole := models.Role{
		Id:   role.Id,
		Name: "deleted:" + role.Name + ":" + strings.Split(role.Id, "-")[0],
		Slug: "deleted:" + role.Slug + ":" + strings.Split(role.Id, "-")[0],
	}

	if err := r.rr.Delete(&deleteRole); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while deleting user",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Role deleted",
	})
}
