package middleware

import (
	"go-gin-clean-architecture/app/helpers"
	"go-gin-clean-architecture/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsUser(jwtService services.JWTservice) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO
		cookieToken, err := ctx.Cookie("token")
		if err != nil {
			res := helpers.BuildErrorResponse("Token Not Found1", err.Error(), nil)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		token, err := jwtService.ValidateToken(cookieToken)
		if err != nil {
			res := helpers.BuildErrorResponse("Token Not Valid midd1", err.Error(), nil)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		if token.Valid {
			ctx.Next()
		} else {
			res := helpers.BuildErrorResponse("Token Not Valid midd2", "Token Not Valid", nil)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

	}
}
