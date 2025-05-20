package handlers

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

// TODO: Implement the get notifications by user ID handler
func GetNotificationsByUserIDHandler(ctx *fasthttp.RequestCtx) {
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
		// Get the user ID from the path
		userID, err := strconv.Atoi(string(ctx.UserValue("user-id").(string)))

		if err != nil {
			log.Printf("Error converting user ID to int at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error converting user ID to int"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		if userID == 0 {
			log.Printf("User ID is not provided at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User ID is not provided"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		// Get the notifications
		notifications := crud.GetNotificationsByRecieverID(userID)

		if err := json.NewEncoder(ctx).Encode(response.NotificationsResp{
			Status:        response.ResponseMessage{Code: "0"},
			Notifications: notifications,
		}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error encoding response"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}
	}
}
