package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/argon"
	"iecare-api/src/app/pkg/paginate"
	"iecare-api/src/app/services"
	"iecare-api/src/app/utils"
	"iecare-api/src/app/validators"
	"strconv"
	"strings"
)

// UsersController is the controller for users
type UsersController struct {
	us services.IUserServices
}

// NewUsersController creates a new instance of the user controller
func NewUsersController(us services.IUserServices) *UsersController {
	return &UsersController{us}
}

func (u *UsersController) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")

	users, err := u.us.List(paginate.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Search:      search,
		Sort:        sort,
		Order:       order,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting users",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(users)
}

func (u *UsersController) Get(c *fiber.Ctx) error {
	uuid := c.Params("userId")

	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	user, err := u.us.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	return c.JSON(user.PublicUser())
}

func (u *UsersController) Store(c *fiber.Ctx) error {
	data := models.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	user := models.User{
		Role: models.RoleUser,
	}
	if err := mergo.Merge(&user, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(user); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	newUser, err := u.us.Store(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while creating user",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(newUser.PublicUser())
}

func (u *UsersController) Edit(c *fiber.Ctx) error {
	uuid := c.Params("userId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	data := models.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	user, err := u.us.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	dstUser := models.User{
		Id: user.Id,
	}
	if err := mergo.Merge(&dstUser, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidatePartialStruct(dstUser); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	editedUser, err := u.us.Edit(&dstUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while updating user",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(editedUser.PublicUser())
}

func (u *UsersController) Delete(c *fiber.Ctx) error {
	uuid := c.Params("userId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	user, err := u.us.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
			"status":  fiber.StatusNotFound,
			"display": true,
		})
	}

	deleteUser := models.User{
		Id:       user.Id,
		Email:    "deleted:" + user.Email + ":" + strings.Split(user.Id, "-")[0],
		UserName: "deleted:" + user.UserName + ":" + strings.Split(user.Id, "-")[0],
	}

	if err := u.us.Delete(&deleteUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while deleting user",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func (u *UsersController) SignIn(c *fiber.Ctx) error {
	data := models.Login{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(data); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	user, err := u.us.FindByMany([]string{"email", "user_name"}, data.Uid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"status":  fiber.StatusNotFound,
			"display": true,
			"error":   err.Error(),
		})
	}

	if match, _ := argon.ComparePasswordAndHash(data.Password, user.Password); match == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
			"status":  fiber.StatusUnauthorized,
			"display": true,
		})
	}

	token, err := utils.GenerateJwt(user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while generating token",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"status":  fiber.StatusOK,
		"display": false,
		"user":    user.PublicUser(),
		"token":   token,
	})
}

func (u *UsersController) SignUp(c *fiber.Ctx) error {
	data := models.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	user := models.User{}
	if err := mergo.Merge(&user, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(user); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	newUser, err := u.us.Store(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while creating user",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	token, err := utils.GenerateJwt(newUser.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while generating token",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "SignUp successful. Please login to continue",
		"status":  fiber.StatusOK,
		"display": false,
		"user":    newUser.PublicUser(),
		"token":   token,
	})
}
