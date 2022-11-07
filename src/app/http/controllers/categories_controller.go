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

type CategoriesController struct {
	cs services.ICategoryServices
}

func NewCategoriesController(cs services.ICategoryServices) *CategoriesController {
	return &CategoriesController{cs}
}

func (s *CategoriesController) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")

	categories, err := s.cs.List(paginate.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Search:      search,
		Sort:        sort,
		Order:       order,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting categories",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(categories)
}

func (s *CategoriesController) Get(c *fiber.Ctx) error {
	uuid := c.Params("categoryId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	category, err := s.cs.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	return c.JSON(category.PublicCategory())
}

func (s *CategoriesController) Store(c *fiber.Ctx) error {
	data := models.Category{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	category := models.Category{}
	if err := mergo.Merge(&category, data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	if errors := validators.ValidateStruct(&category); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  errors,
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	newCategory, err := s.cs.Store(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while storing category",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(newCategory.PublicCategory())
}

func (s *CategoriesController) Edit(c *fiber.Ctx) error {
	uuid := c.Params("categoryId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	data := models.Category{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	category, err := s.cs.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	dstCategory := models.Category{
		Id: category.Id,
	}
	if err := mergo.Merge(&dstCategory, data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	if errors := validators.ValidateStruct(dstCategory); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  errors,
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	editedCategory, err := s.cs.Edit(&dstCategory)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while editing category",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(editedCategory.PublicCategory())
}

func (s *CategoriesController) Delete(c *fiber.Ctx) error {
	uuid := c.Params("categoryId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	category, err := s.cs.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	deleteCategory := models.Category{
		Id:   category.Id,
		Name: "deleted:" + category.Name + ":" + strings.Split(category.Id, "-")[0],
	}

	if err := s.cs.Delete(&deleteCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while deleting category",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
