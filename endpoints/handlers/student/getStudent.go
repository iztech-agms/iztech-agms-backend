package student

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/endpoints/response"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

type GetStudentsByAdvisor struct {
	Response response.ResponseMessage
	Users    []crud.User
}

// User ID al, ilişikli studentları bul
func GetStudentListByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		uid, err := strconv.Atoi(ctx.UserValue("id").(string))
		if err != nil {
			log.Printf("Error converting user ID to int at endpoint (%s): %v", path, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error converting user ID to int"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		if uid == 0 {
			log.Printf("User ID is not provided at endpoint (%s).", path)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "User ID is not provided"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		var students []crud.Student
		role := crud.GetUserByID(uid).Role
		if role == "advisor" {
			students = crud.GetStudentsByAdvisorID(uid)
		} else if role == "department_secretary" {
			students = crud.GetStudentsByDepartmentSecretaryID(uid)
		} else if role == "faculty_secretary" {
			students = crud.GetStudentsByFacultySecretaryID(uid)
		} else if role == "student_affairs" {
			students = crud.GetStudents()
		}

		if len(students) == 0 {
			log.Printf("No match for students at endpoint (%s).", path)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No match for students"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		var userList []crud.User
		for _, student := range students {
			s := crud.GetUserByID(student.ID)
			userList = append(userList, s)
		}
		resp := GetStudentsByAdvisor{
			Users:    userList,
			Response: response.ResponseMessage{Code: "0", Message: "Success"},
		}
		if err = json.NewEncoder(ctx).Encode(resp); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			return
		}
	}
}
