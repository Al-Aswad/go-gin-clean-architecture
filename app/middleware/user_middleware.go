package middleware

import (
	"go-gin-note-app/app/helpers"
	"go-gin-note-app/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

		if token.Valid {
			ctx.Next()
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// Token Note Valid
				res := helpers.BuildErrorResponse("Token tidak valid", err.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token Expire
				res := helpers.BuildErrorResponse("Token Expire", err.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			} else {
				res := helpers.BuildErrorResponse("Token tidak valid", err.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			}
		} else {
			res := helpers.BuildErrorResponse("Token tidak valid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
}
