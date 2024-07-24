package channelhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/database"
	"github.com/masterghost2002/videotube/types"
)

func DeleteChannel(c *fiber.Ctx) error {
	channelIdString := c.Params("id")

	loggedUser, ok := c.Locals("logged_user").(database.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(&types.Response{
			Error:       true,
			ErrorFields: ok,
			Data:        nil,
			Message:     "Unauthorized",
		})
	}
	channelId, convertErr := strconv.ParseInt(channelIdString, 10, 64)
	if convertErr != nil {
		return c.Status(400).JSON(&types.Response{
			Error:       true,
			ErrorFields: convertErr,
			Message:     "Invalid channel id",
			Data:        nil,
		})
	}
	// get channel first then check is the logged user trying to delete the channel

	channel, getChannelErr := database.Storage.GetChannelByUserId(context.Background(), loggedUser.ID)
	if getChannelErr != nil {
		return getChannelErr
	}

	if channel.ID != channelId {
		return c.Status(400).JSON(&types.Response{
			Error:       true,
			ErrorFields: nil,
			Message:     "This channel doesn't belongs to you",
			Data:        nil,
		})
	}

	result := database.Storage.DeleteChannel(context.Background(), channelId)

	return c.Status(200).JSON(&types.Response{
		Error:       false,
		ErrorFields: nil,
		Message:     "Channel Deleted",
		Data:        result,
	})
}
