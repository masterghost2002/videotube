package channelhandlers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/database"
	"github.com/masterghost2002/videotube/internals/validations"
	"github.com/masterghost2002/videotube/types"
)

func CreateChannel(c *fiber.Ctx) error {
	var channelData validations.CreateChannelFormData
	user, ok := c.Locals("logged_user").(database.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	if err := c.BodyParser(&channelData); err != nil {
		return err
	}

	if err := validations.Validate.Struct(channelData); err != nil {
		return err
	}

	// check if the channel already exist
	if _, err := database.Storage.GetChannelById(context.Background(), user.ID); err != sql.ErrNoRows {
		return errors.New("channel already exists")
	}

	createChannelParams := database.CreateChannelParams{
		Name:   channelData.Name,
		UserID: user.ID,
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
