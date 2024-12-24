package config

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

func maxFileSizeValidation(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	// Dapatkan parameter maxSize
	maxSize, err := strconv.ParseInt(fl.Param(), 10, 64)
	if err != nil {
		return false
	}

	// Validasi ukuran file
	return file.Size <= maxSize*1024*1024
}

func InitializeValidator() *validator.Validate {
	var validatorInstance = validator.New()
	// Tambahkan validasi kustom untuk 'datetime'
	validatorInstance.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		format := "2006-01-02"
		_, err := time.Parse(format, fl.Field().String())
		return err == nil
	})
	validatorInstance.RegisterValidation("maxSize", maxFileSizeValidation)
	validatorInstance.RegisterValidation("extensionFile", validateFileExtensionValidation)
	validatorInstance.RegisterValidation("obligatoryFile", requiredFileValidationValidation)
	return validatorInstance
}

func requiredFileValidationValidation(fl validator.FieldLevel) bool {
	_, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}
	return true
}

// Validasi ekstensi file
func validateFileExtensionValidation(fl validator.FieldLevel) bool {
	// Ambil file dari field
	file, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	// Ambil parameter ekstensi yang diperbolehkan
	allowedExtension := fl.Param() // Misalnya "png,jpg,jpeg"

	// Pisahkan parameter ekstensi menjadi slice
	extSlice := strings.Split(allowedExtension, " ")

	// Ambil ekstensi file
	filename := strings.ToLower(file.Filename)
	for _, ext := range extSlice {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}
