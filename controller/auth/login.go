package auth

import (
	"database/sql"

	"aphro.web/database"
	"aphro.web/model"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	// user form data from request
	user := new(model.User)
	// user data from database
	var regUser *model.User

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	db := database.GetDbInstance()
	// fetch user data from database using email from the request
	regUser, err := db.FetchUser(user.Email)

	if err != nil {
		// unable to find any entry in database
		if err == sql.ErrNoRows {
			return ctx.JSON(
				fiber.Map{"error": "Unable to find a user with the specified data"})
		}
		return err // to be handled by default error handler
	}

	// checking request password and hashed password from database
	if err := CheckPasswordHash(user.Password, regUser.Password); err != nil {
		return ctx.JSON(
			fiber.Map{"error": "password and email mismatched"})
	}

	// create jwt auth taken i.e. Authorization: Bearer <TOKEN>
	token, err := createJWTToken(regUser)
	if err != nil {
		return ctx.JSON(
			fiber.Map{"error": "unable to create auth token " + err.Error()})
	}

	return ctx.JSON(fiber.Map{"jwt": token})
}
