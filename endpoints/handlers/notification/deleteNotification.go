package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
	"strconv"
)

// TODO: Implement the delete notification handler
func DeleteNotificationHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())
	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
	default:
		nid, err := strconv.Atoi(ctx.UserValue("id").(string))
		if err != nil {
			log.Printf("Error parsing notification id %s: %v", ctx.UserValue("id").(string), err)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Message: "Error parsing notification id"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		if nid == 0 {
			log.Printf("Notification id is not provided")
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Notification id is not provided"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		err = crud.DeleteNotificationByID(nid)
		if err != nil {
			log.Printf("Error deleting notification %s: %v", path, err)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Message: "Error deleting notification"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Message: "Notification deleted successfully"}); err != nil {
		}
	}

}

