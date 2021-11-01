package auth

import (
	"strings"

	"aphro.web/database"
	"aphro.web/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(ctx *fiber.Ctx) error {
	user := new(model.User)
	user.Id = uuid.New()

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = database.GetDbInstance().InsertUser(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ctx.JSON(fiber.Map{"error": "Duplicated key"})
		}
		return err
	}

	return ctx.JSON(user)
}
