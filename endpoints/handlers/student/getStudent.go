package student

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/crud/customized"
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

func GetStudentDetailedByIDHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return

	default:
		studentID, err := strconv.Atoi(ctx.UserValue("id").(string))
		if err != nil {
			log.Printf("Error converting student ID to int at endpoint (%s): %v", path, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Error converting student ID to int"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		if studentID == 0 {
			log.Printf("Student ID is not provided at endpoint (%s).", path)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Student ID is not provided"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		studentDetailed := customized.GetStudentListDetailedByUserIDs([]int{studentID})
		if len(studentDetailed) == 0 {
			log.Printf("Student %d not found", studentID)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Student not found"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		if len(studentDetailed) != 1 {
			log.Printf("Multiple students found for ID %d", studentID)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Multiple students found"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.StudentsDetailedResp{
			Status:   response.ResponseMessage{Code: "0"},
			Students: studentDetailed,
		}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			return
		}
	}
}

func GetAllUsersTest(ctx *fasthttp.RequestCtx) {
	resp := crud.GetUsers()
	if err := json.NewEncoder(ctx).Encode(resp); err != nil {
		log.Printf("Error encoding response at endpoint Test Users: %v", err)
		return
	}

}
