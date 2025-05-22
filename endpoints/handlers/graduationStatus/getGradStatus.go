package gradStatus

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

func GetGradStatusByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		// Get student ID from path
		stdId, err := strconv.Atoi(ctx.UserValue("id").(string))
		if err != nil {
			log.Printf("Error converting user ID to int at endpoint (%s): %v", path, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error converting user ID to int"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		if stdId == 0 {
			log.Printf("User ID is not provided at endpoint (%s).", path)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User ID is not provided"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Get graduation status matching with stdId
		res := crud.GetGraduationStatusByStudentID(stdId)
		resp := make([]crud.GraduationStatus, 1)
		resp[0] = res
		if err := json.NewEncoder(ctx).Encode(response.GraduationStatusResp{
			Status:           response.ResponseMessage{Code: "0"},
			GraduationStatus: resp,
		}); err != nil {

		}
	}
}

func GetGradStatusByGradYear(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())
	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		year, err := strconv.Atoi(ctx.UserValue("year").(string))
		if err != nil {
			log.Printf("Error converting graduation year to int at endpoint (%s): %v", path, err)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error converting graduation year to int"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		if year == 0 {
			log.Printf("Graduation year is not provided at endpoint (%s).", path)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		res := crud.GetGraduationStatusesByYear(year)
		if err = json.NewEncoder(ctx).Encode(response.GraduationStatusResp{
			Status:           response.ResponseMessage{Code: "0"},
			GraduationStatus: res,
		}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			return
		}
	}
}
