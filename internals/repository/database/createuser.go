package repository

import (
	"github.com/masterghost2002/videotube/internals/models"
	"github.com/masterghost2002/videotube/internals/registry"
)

func CreateUser(data models.User) error {
	storage := registry.GetDB()

	user := models.User(data)

	result := storage.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
