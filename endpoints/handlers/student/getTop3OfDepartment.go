package student

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/crud/customized"
	"graduation-system/endpoints/response"
	"graduation-system/util/studentUtil"
	"log"
	"strings"

	"github.com/valyala/fasthttp"
)

func GetTop3OfStudentsOfDepartmentHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		departmentName := string(ctx.UserValue("department-name").(string))
		// Replace _ to space
		departmentName = strings.ReplaceAll(departmentName, "_", " ")

		log.Printf("Department name: %s", departmentName)
		studentIDs := crud.GetStudentIDsByDepartmentName(departmentName)
		if len(studentIDs) == 0 {
			log.Printf("(Error) : no students found for department name: %s", departmentName)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No students found for department name"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		studentsDetailed := customized.GetStudentListDetailedByUserIDs(studentIDs)
		studentsDetailed = studentUtil.FilterStudentListByAll(studentsDetailed)
		if len(studentsDetailed) == 0 {
			log.Printf("(Error) : no students found for department name: %s", departmentName)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No students found for department name"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		// Sort
		studentUtil.StudentSort(studentsDetailed)

		// Get top 3 students
		top3Students := studentsDetailed[:3]

		// Return top 3 students
		if err := json.NewEncoder(ctx).Encode(top3Students); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}

	}
}
