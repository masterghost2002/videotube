package channelhandlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/database"
	"github.com/masterghost2002/videotube/types"
)

func GetChannels(c *fiber.Ctx) error {
	channels, err := database.Storage.GetChannels(context.Background())
	if err != nil {
		return err
	}

	return c.Status(200).JSON(&types.Response{
		Error:       false,
		Message:     "Channels",
		ErrorFields: nil,
		Data:        channels,
	})

}
