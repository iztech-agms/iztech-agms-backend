package student

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/crud/customized"
	"graduation-system/endpoints/response"
	"graduation-system/util/studentUtil"
	"log"

	"github.com/valyala/fasthttp"
)

func GetTop3StudentsOfFacultyHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		facultyName := string(ctx.UserValue("faculty-name").(string))

		departments := crud.GetDepartmentByFacultyName(facultyName)
		if len(departments) == 0 { // If no departments found for faculty name, return error
			log.Printf("(Error) : no departments found for faculty name: %s", facultyName)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No departments found for faculty name"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		studentIDs := []int{}
		for _, department := range departments {
			advisors := crud.GetAdvisorsByDepartmentName(department.Name)
			for _, advisor := range advisors {
				studentIDs = append(studentIDs, customized.GetStudentIDsByUserID(advisor.ID)...)
			}
		}

		// Get top 3 students by student IDs
		studentsDetailed := customized.GetStudentListDetailedByUserIDs(studentIDs)
		studentsDetailed = studentUtil.FilterStudentListByAll(studentsDetailed)
		if len(studentsDetailed) == 0 {
			log.Printf("(Error) : no students found for faculty name: %s", facultyName)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No students found for faculty name"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Sort students by graduation status
		studentUtil.StudentSort(studentsDetailed)

		// Get top 3 students
		top3Students := studentsDetailed[:3]

		// Return top 3 students
		if err := json.NewEncoder(ctx).Encode(top3Students); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
	}
}
