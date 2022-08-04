package middleware

import (
	"fmt"
	"gin-note-app/helpers"
	"gin-note-app/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsUser(jwtService services.JWTservice) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieToken := ctx.Request.Header["Authorization"]
		fmt.Println("cookie1 ===========", cookieToken)
		if cookieToken == nil {
			res := helpers.BuildErrorResponse("Token Not Found1", "Token Kosong !", nil)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		tokenString := strings.Replace(cookieToken[0], "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(tokenString)

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
