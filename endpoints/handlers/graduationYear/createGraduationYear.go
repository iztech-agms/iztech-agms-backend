package graduationYear

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func CreateGraduationYearHandler(ctx *fasthttp.RequestCtx) {
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

		// Check if the graduation year already exists
		graduationYear := crud.GetGraduationYearByYear(currentYear)
		if graduationYear.Year != 0 {
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "1", Message: "Graduation year already exists"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		// If not exists, create the graduation year
		gradYearStruct := crud.GraduationYear{
			Year:      currentYear,
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 14).Truncate(24 * time.Hour).Add(6 * time.Hour),
		}

		if err := crud.CreateGraduationYear(&gradYearStruct); err != nil {
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "1", Message: "Error creating graduation year"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		// Bu nokta
		// Bu noktada tüm ilgili hocalara notification gönderilecek.
		// Create notification for all advisors
		for _, advisor := range crud.GetAdvisors() {
			advisorNotification := crud.Notification{
				UserID:             advisor.ID,
				Message:            "Student affairs has requested graduation lists",
				Title:              "Graduation List Request",
				IsNotificationRead: false,
			}
			if err := crud.CreateNotification(&advisorNotification); err != nil {
				log.Printf("Error creating notification for advisor %d: %v", advisor.ID, err)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "500", Message: "Error creating notification for advisor"}); err != nil {
					log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
				}
				return
			}
		}

		// Tüm 0 olan graduation system confirm status'ler 2 ye çevrilcek.
		for _, grad_status := range crud.GetGraduationStatuses() {
			if grad_status.IsSystemConfirmed == 0 {
				grad_status.IsSystemConfirmed = 2
				crud.UpdateGraduationStatus(grad_status)
			}
		}

		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
		}
		return
	}
}
