package auth

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"graduation-system/util"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

type BasicLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status response.ResponseMessage `json:"status"`
	User   crud.User                `json:"user"`
	Token  string                   `json:"token"`
}

// AuthLoginHandler handles the login request for the user
func AuthLoginHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
		}
		return

	default:
		var loginReq BasicLoginReq
		if err := json.Unmarshal(ctx.PostBody(), &loginReq); err != nil {
			log.Printf("Error decoding request at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error decoding request"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		user := crud.GetUserByUsername(loginReq.Username)
		if user.ID == 0 {
			log.Printf("User not found at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User not found"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		if util.CheckPasswordHash(loginReq.Password, user.Password) {
			log.Printf("Invalid password at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Wrong password"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			log.Printf("Error generating JWT token at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error generating JWT token"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		loginResponse := LoginResponse{
			Status: response.ResponseMessage{Code: "0", Message: "Login successful"},
			User:   user,
			Token:  tokenString,
		}

		if err := json.NewEncoder(ctx).Encode(loginResponse); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			return
		}
	}
}
