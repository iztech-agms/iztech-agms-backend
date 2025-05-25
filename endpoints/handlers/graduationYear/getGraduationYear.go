package graduationYear

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func GetGraduationYearHandler(ctx *fasthttp.RequestCtx) {
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
		// Get the current year (2025)
		currentYear := time.Now().Year()

		// Get the graduation year by year
		graduationYear := crud.GetGraduationYearByYear(currentYear)
		if graduationYear.Year == 0 {
			if err := json.NewEncoder(ctx).Encode(response.GraduationYearResp{
				Status:   response.ResponseMessage{Code: "1", Message: ""},
				GradYear: []crud.GraduationYear{},
			}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		// Return the graduation year
		if err := json.NewEncoder(ctx).Encode(response.GraduationYearResp{
			Status:   response.ResponseMessage{Code: "0", Message: ""},
			GradYear: []crud.GraduationYear{graduationYear},
		}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
		}
	}
}
