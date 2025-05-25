package graduationYear

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"

	"github.com/valyala/fasthttp"
)

func UpdateGraduationYearHandler(ctx *fasthttp.RequestCtx) {
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
		var gradYearStruct crud.GraduationYear
		if err := json.Unmarshal(ctx.Request.Body(), &gradYearStruct); err != nil {
			log.Printf("Error decoding request body at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Message: "Error decoding request body"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		// Update the graduation year
		if err := crud.UpdateGraduationYear(gradYearStruct); err != nil {
			log.Printf("Error updating graduation year at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "5", Message: "Error updating graduation year"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Message: "Graduation year updated successfully"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
		}
	}
}
