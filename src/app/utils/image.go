package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/h2non/bimg"
	"io"
	"mime/multipart"
	"os"
	"regexp"
	"strings"
)

func IsExistFile(filename string, path string) bool {
	if _, err := os.Stat(path + filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteFile(filename string, path string) error {
	if err := os.Remove(path + filename); err != nil {
		return err
	}
	return nil
}

func IsExistFolder(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

func ImageCompression(fileHeader *multipart.FileHeader, quality int) (string, error) {
	filename := strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(strings.Split(fileHeader.Filename, ".")[0], "")+"_"+utils.UUIDv4()) + ".webp"

	file, err := fileHeader.Open()
	if err != nil {
		return filename, err
	}

	buffer, err := io.ReadAll(file)
	if err != nil {
		return filename, err
	}

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return filename, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
	if err != nil {
		return filename, err
	}

	if err := bimg.Write(fmt.Sprint("public/uploads/"+filename), processed); err != nil {
		return filename, err
	}

	return filename, nil
}
