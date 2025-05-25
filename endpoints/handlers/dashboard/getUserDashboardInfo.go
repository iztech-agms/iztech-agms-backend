package dashboard

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

type UserDashboardInfo struct {
	DepartmentName string `json:"department_name"`
	OfficeLocation string `json:"office_location"`
}

type UserDashboardInfoResponse struct {
	Status response.ResponseMessage `json:"status"`
	Data   UserDashboardInfo        `json:"data"`
}

// Get user dashboard info
func GetUserDashboardInfoHandler(ctx *fasthttp.RequestCtx) {
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
		var userDashboardInfo UserDashboardInfo
		userID, err := strconv.Atoi(ctx.UserValue("user-id").(string))
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

		// Get user
		user := crud.GetUserByID(userID)
		if user.ID == 0 {
			log.Printf("User not found at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User not found"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		switch user.Role {
		case "student_affairs":
			stdAffairs := crud.GetStudentAffairsByID(userID)
			if stdAffairs.ID == 0 {
				log.Printf("Student affairs not found at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Student affairs not found"}); err != nil {
					log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
				}
				return
			}

			userDashboardInfo = UserDashboardInfo{
				OfficeLocation: stdAffairs.OfficeLocation,
			}

		case "advisor":
			advisor := crud.GetAdvisorByID(userID)
			if advisor.ID == 0 {
				log.Printf("Advisor not found at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Advisor not found"}); err != nil {
					log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
				}
				return
			}

			userDashboardInfo = UserDashboardInfo{
				OfficeLocation: advisor.OfficeLocation,
				DepartmentName: advisor.DepartmentName,
			}

		case "department_secretary":
			departmentSecretary := crud.GetDepartmentSecretaryByID(userID)
			if departmentSecretary.ID == 0 {
				log.Printf("Department secretary not found at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Department secretary not found"}); err != nil {
					log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
				}
				return
			}

			userDashboardInfo = UserDashboardInfo{
				OfficeLocation: departmentSecretary.OfficeLocation,
				DepartmentName: departmentSecretary.DepartmentName,
			}

		case "faculty_secretary":
			facultySecretary := crud.GetFacultySecretaryByID(userID)
			if facultySecretary.ID == 0 {
				log.Printf("Faculty secretary not found at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Faculty secretary not found"}); err != nil {
					log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
				}
				return
			}

			userDashboardInfo = UserDashboardInfo{
				OfficeLocation: facultySecretary.OfficeLocation,
				DepartmentName: facultySecretary.FacultyName,
			}

		default:
			log.Printf("User role not found at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User role not found"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(UserDashboardInfoResponse{
			Status: response.ResponseMessage{Code: "0"},
			Data:   userDashboardInfo,
		}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
		}
	}
}
