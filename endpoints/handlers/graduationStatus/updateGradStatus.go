package gradStatus

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
)

func UpdateGradStatus(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		var body crud.GraduationStatus
		if err := json.Unmarshal(ctx.PostBody(), &body); err != nil {
			log.Printf("Error decoding request at endpoint (%s): %v", path, err)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error decoding request"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		if err := crud.UpdateGraduationStatus(body); err != nil {
			log.Printf("Internal Server Error : %v", err)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Internal Server Error"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Message: "Success"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			return
		}
	}
}

