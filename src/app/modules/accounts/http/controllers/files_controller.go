package controllers

import (
	"github.com/gofiber/fiber/v2"
	uuid "github.com/gofiber/fiber/v2/utils"
	"iecare-api/src/app/shared/utils"
	"strings"
)

type File struct {
	FileName         string `json:"filename"`
	OriginalFileName string `json:"original_filename"`
	FileFormat       string `json:"format"`
	FileType         string `json:"type"`
	Size             int64  `json:"size"`
	Url              string `json:"url"`
}

func Store(c *fiber.Ctx) error {
	// Create a folder to upload if not exists
	if !utils.IsExistFolder("public/uploads/") {
		if err := utils.CreateFolder("public/uploads/"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error while creating folder",
				"error":   err.Error(),
				"status":  fiber.StatusInternalServerError,
				"display": true,
			})
		}
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting form",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})

	}

	files := form.File["file"]
	var links []*File

	for _, file := range files {
		var link File
		var filename string

		format := strings.Split(file.Header["Content-Type"][0], "/")[1]
		fileType := strings.Split(file.Header["Content-Type"][0], "/")[0]
		if fileType == "image" {
			name, err := utils.ImageCompression(file, 80)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Error while opening file",
					"error":   err.Error(),
					"status":  fiber.StatusBadRequest,
					"display": true,
				})
			}
			filename = name
		} else {
			filename = uuid.UUIDv4()
			if err := c.SaveFile(file, "public/uploads/"+filename+"."+format); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error while getting form",
					"error":   err.Error(),
					"status":  fiber.StatusInternalServerError,
					"display": true,
				})
			}
		}

		link.FileName = filename
		link.OriginalFileName = strings.Split(file.Filename, ".")[0]
		link.Size = file.Size
		link.Url = c.BaseURL() + "/files/uploads/" + filename
		link.FileFormat = format
		link.FileType = fileType

		links = append(links, &link)
	}

	return c.JSON(links)
}
