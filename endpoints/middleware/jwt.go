package middleware

import (
	"encoding/json"
	"graduation-system/endpoints/response"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

func JWTMiddleware(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenString := string(ctx.Request.Header.Peek("Authorization"))
		if tokenString == "" {
			return
		}
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			log.Printf("Error parsing JWT token at JWT Middleware: %v", err)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error parsing JWT token"}); err != nil {
				log.Printf("Error encoding response at JWT Middleware: %v", err)
			}
			return
		}
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("Error parsing JWT token at JWT Middleware: %v", err)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error parsing JWT token"}); err != nil {
				log.Printf("Error encoding response at JWT Middleware: %v", err)
			}
			return
		}
		// Check if the expiry time gone or not
		expTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().Unix() > expTime.Unix() {
			log.Printf("Token expired at JWT Middleware")
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Message: "Token expired at endpoint"}); err != nil {
				log.Printf("Error encoding response at JWT Middleware: %v", err)
			}
			return
		}

		// Create new token
		uname := claims["username"].(string)
		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": uname,
			"exp":      time.Now().Add(time.Minute * 1).Unix(),
		})
		tokenString, err = token.SignedString([]byte("secret"))
		if err != nil {
			log.Printf("Error generating JWT token at JWT Middleware : %v", err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error generating JWT token"}); err != nil {
				log.Printf("Error encoding response at JWT Middleware : %v", err)
			}
			return
		}
		ctx.Request.Header.Set("Authorization", tokenString)
		requestHandler(ctx)
	}
}
