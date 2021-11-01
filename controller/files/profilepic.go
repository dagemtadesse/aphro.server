package files

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func PostProfilePic(ctx *fiber.Ctx) error {

	file, err := ctx.FormFile("image")

	if err != nil {

		if errors.Is(err, fasthttp.ErrMissingFile) {
			return ctx.JSON(fiber.Map{"error": "image file mising"})
		}

		return err
	}

	authClaim := ctx.Locals("authClaims")
	if authClaim == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "anuthroized request"})
	}

	authMap := authClaim.(jwt.MapClaims)
	if authMap["authorized"] == true {
		// path := fmt.Sprintf("")
		ctx.SaveFile(file, file.Filename)
		return ctx.SendString("Saving file")
	}

	return ctx.SendString("Not Implemented")
}
