package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
	"iecare-api/src/app/services"
	"iecare-api/src/app/validators"
	"os"
	"strconv"
	"strings"
)

type ProvidersController struct {
	pr services.IProviderServices
}

func NewProvidersController(pr services.IProviderServices) *ProvidersController {
	return &ProvidersController{pr}
}

func (p *ProvidersController) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")

	providers, err := p.pr.List(paginate.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Search:      search,
		Sort:        sort,
		Order:       order,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting providers",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(providers)
}

func (p *ProvidersController) Get(c *fiber.Ctx) error {
	uuid := c.Params("providerId")

	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	provider, err := p.pr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Provider not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	return c.JSON(provider.PublicProvider())
}

func (p *ProvidersController) Store(c *fiber.Ctx) error {
	data := models.Provider{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing request body",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	userId := c.Locals("userId").(string)
	fmt.Fprintf(os.Stdout, "userId: %s", userId)
	provider := models.Provider{
		UserId: userId,
	}
	if err := mergo.Merge(&provider, data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while merging request body",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	if errors := validators.ValidateStruct(provider); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation errors",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	newProvider, err := p.pr.Store(&provider)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while creating provider",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": true,
		})
	}

	return c.JSON(newProvider.PublicProvider())
}

func (p *ProvidersController) Edit(c *fiber.Ctx) error {
	uuid := c.Params("providerId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	data := models.Provider{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing request body",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	provider, err := p.pr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Provider not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	dstProvider := models.Provider{
		Id: provider.Id,
	}
	if err := mergo.Merge(&dstProvider, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging request body",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(dstProvider); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation errors",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	provider, err = p.pr.Edit(&dstProvider)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while updating provider",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(provider.PublicProvider())
}

func (p *ProvidersController) Delete(c *fiber.Ctx) error {
	uuid := c.Params("providerId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	provider, err := p.pr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Provider not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	deleteProvider := models.Provider{
		Id:    provider.Id,
		Name:  "deleted:" + provider.Name + strings.Split(provider.Id, "-")[0],
		Email: "deleted:" + provider.Email + ":" + strings.Split(provider.Id, "-")[0],
		Phone: "deleted:" + provider.Phone + ":" + strings.Split(provider.Id, "-")[0],
	}
	if err := p.pr.Delete(&deleteProvider); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while deleting provider",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Provider deleted successfully",
	})
}
