package student

import (
	"encoding/json"
	"fmt"
	"graduation-system/crud"
	"graduation-system/crud/customized"
	"graduation-system/endpoints/response"
	"graduation-system/util/studentUtil"
	"log"
	"strings"

	"github.com/valyala/fasthttp"
)

func ApplyForGraduationHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())
	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}

		return
	default:
		var studentID int
		if err := json.Unmarshal(ctx.Request.Body(), &studentID); err != nil {
			log.Printf("Error decoding request body at endpoint (%s): %v", path, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Message: "Error decoding request body"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Get student
		students := customized.GetStudentListDetailedByUserIDs([]int{studentID})
		if len(students) == 0 {
			log.Printf("Student %d not found", studentID)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "100", Message: "Student not found"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		studentDetailed := students[0]

		// Check if student can graduate
		if code, values := studentUtil.CanStudentApply(studentID); code == "0" {
			log.Printf("Student %d cannot graduate", studentID)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: code, Message: strings.Join(values, ", ")}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Update graduation status of student
		studentDetailed.GraduationStatus.IsSystemConfirmed = 3 // If student can graduate system approves
		studentDetailed.GraduationStatus.IsAdvisorConfirmed = 2 // The rest remain pending
		studentDetailed.GraduationStatus.IsDepSecConfirmed = 2
		studentDetailed.GraduationStatus.IsFacultyConfirmed = 2
		studentDetailed.GraduationStatus.IsStdAffConfirmed = 2
		if err := crud.UpdateGraduationStatus(studentDetailed.GraduationStatus); err != nil {
			log.Printf("Error updating graduation status of student %d: %v", studentID, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "500", Message: "Error updating graduation status"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Create notification for advisor
		advisorNotification := crud.Notification{
			UserID:             studentDetailed.AdvisorID,
			Message:            fmt.Sprintf("Student %s %s has applied for graduation", studentDetailed.User.FirstName, studentDetailed.User.LastName),
			Title:              "Graduation Application",
			IsNotificationRead: false,
		}
		if err := crud.CreateNotification(&advisorNotification); err != nil {
			log.Printf("Error creating notification for advisor %d: %v", studentDetailed.AdvisorID, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "500", Message: "Error creating notification for advisor"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Encode response
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
	}
}
