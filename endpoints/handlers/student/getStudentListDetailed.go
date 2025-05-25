package student

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/crud/customized"
	"graduation-system/endpoints/response"
	"graduation-system/util/studentUtil"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

type GetStudentListDetailedByUserID struct {
	Response response.ResponseMessage
	Students []crud.Student
}

func GetStudentListDetailedByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		userID, err := strconv.Atoi(ctx.UserValue("userid").(string))
		if err != nil {
			log.Printf("Error converting user ID to int at endpoint (%s): %v", path, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error converting user ID to int"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		if userID == 0 {
			log.Printf("User ID is 0 at endpoint (%s).", path)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User ID is 0"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Get user first
		user := crud.GetUserByID(userID)
		if user.ID == 0 {
			log.Printf("User not found at endpoint (%s).", path)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User not found"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		studentIDs := customized.GetStudentIDsByUserID(userID)

		studentsDetailed := customized.GetStudentListDetailedByUserIDs(studentIDs)

		//Filter
		if user.Role == "advisor" {
			studentsDetailed = studentUtil.FilterStudentListByAdvisor(studentsDetailed)
		} else if user.Role == "department_secretary" {
			studentsDetailed = studentUtil.FilterStudentListByDepartmentSecretary(studentsDetailed)
		} else if user.Role == "faculty_secretary" {
			studentsDetailed = studentUtil.FilterStudentListByFacultySecretary(studentsDetailed)
		} else if user.Role == "student_affairs" {
			studentsDetailed = studentUtil.FilterStudentListByStudentAffairs(studentsDetailed)
		}

		if len(studentsDetailed) == 0 {
			//log.Printf("User ID is 0 at endpoint (%s).", path)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No students to display"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(studentsDetailed); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
	}
}
