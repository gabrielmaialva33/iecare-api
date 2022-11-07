package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
	"iecare-api/src/app/services"
	"iecare-api/src/app/validators"
	"strconv"
	"strings"
)

type ServicesController struct {
	ss services.IServiceServices
}

func NewServicesController(ss services.IServiceServices) *ServicesController {
	return &ServicesController{ss}
}

func (s *ServicesController) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")

	listServices, err := s.ss.List(paginate.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Search:      search,
		Sort:        sort,
		Order:       order,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting services",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(listServices)
}

func (s *ServicesController) Get(c *fiber.Ctx) error {
	uuid := c.Params("serviceId")

	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	service, err := s.ss.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Service not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	return c.JSON(service.PublicService())
}

func (s *ServicesController) Store(c *fiber.Ctx) error {
	data := models.Service{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	service := models.Service{}
	if err := mergo.Merge(&service, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(service); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  errors,
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	newService, err := s.ss.Store(&service)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while creating service",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(newService.PublicService())
}

func (s *ServicesController) Edit(c *fiber.Ctx) error {
	uuid := c.Params("serviceId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	data := models.Service{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	service, err := s.ss.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Service not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	dstService := models.Service{
		Id: service.Id,
	}
	if err := mergo.Merge(&dstService, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(service); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	editedService, err := s.ss.Edit(&dstService)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while editing service",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(editedService.PublicService())
}

func (s *ServicesController) Delete(c *fiber.Ctx) error {
	uuid := c.Params("serviceId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	service, err := s.ss.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Service not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	deleteService := models.Service{
		Id:   service.Id,
		Name: "deleted:" + service.Name + ":" + strings.Split(service.Id, "-")[0],
	}
	if err := s.ss.Delete(&deleteService); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while deleting service",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Service deleted successfully",
	})
}
