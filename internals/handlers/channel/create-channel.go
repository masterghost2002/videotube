package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/database"
	"github.com/masterghost2002/videotube/internals/validations"
	"github.com/masterghost2002/videotube/types"
)

func CreateChannel(c *fiber.Ctx) error {
	var channelData validations.CreateChannelFormData
	if err := c.BodyParser(&channelData); err != nil {
		return err
	}

	if err := validations.Validate.Struct(channelData); err != nil {
		return err
	}
	createChannelParams := database.CreateChannelParams{
		Name: channelData.Name,
	}
	result, err := database.Storage.CreateChannel(context.Background(), createChannelParams)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(&types.Response{
		Error:       false,
		Message:     "Channel Created",
		ErrorFields: nil,
		Data:        result,
	})
}
