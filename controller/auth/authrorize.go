package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleWare(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("authorization")

	if authHeader != "" {
		tokenString := strings.Split(authHeader, " ")[1]
		payload, err := ParseJWTToken(tokenString)

		if err == nil {
			ctx.Locals("authClaims", payload)
		}

		//debug code
		if ctx.Query("jwtDebug", "false") == "true" {
			return ctx.JSON(payload)
		}

	}

	return ctx.Next()
}
