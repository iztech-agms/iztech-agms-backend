package studentUtil

import "graduation-system/crud"

func CanStudentGraduate(studentID int) bool {
	grad_stat := crud.GetGraduationStatusByStudentID(studentID)
	if grad_stat.StudentGPA >= 2.0 && grad_stat.StudentECTS >= 240 {
		return true
	}

	return false
}
