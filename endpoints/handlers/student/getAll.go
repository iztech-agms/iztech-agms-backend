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

func GetStudentsOfAllHandler(ctx *fasthttp.RequestCtx) {
	var path = string(ctx.Path())

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", path)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "Client canceled the request"}); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}
		return
	default:
		// Get all students
		studentIDs := crud.GetStudentIDs()
		if len(studentIDs) == 0 {
			log.Printf("(Error) : no students found for all.")
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No students found for all"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}

		studentsDetailed := customized.GetStudentListDetailedByUserIDs(studentIDs)
		studentsDetailed = studentUtil.FilterStudentListByAll(studentsDetailed)
		if len(studentsDetailed) == 0 {
			log.Printf("(Error) : no students found for all")
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Message: "No students found for all"}); err != nil {
				log.Printf("Error encoding response at endpoint (%s): %v", path, err)
			}
			return
		}
		// Sort
		studentUtil.StudentSort(studentsDetailed)

		// Return all confirmed students
		if err := json.NewEncoder(ctx).Encode(studentsDetailed); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", path, err)
		}

	}
}
