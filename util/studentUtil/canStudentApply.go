package studentUtil

import "graduation-system/crud"

func CanStudentApply(studentID int) (string, []string) {
	grad_stat := crud.GetGraduationStatusByStudentID(studentID)
	if grad_stat.StudentGPA >= 2.0 && grad_stat.StudentECTS >= 240 && grad_stat.IsSystemConfirmed == 0 {
		return "1", []string{}
	}

	return "0", []string{}
}
