package studentUtil

import "graduation-system/crud/customized"

func FilterStudentListByAll(studentsList []customized.StudentDetailed) []customized.StudentDetailed {
	filteredStudentList := []customized.StudentDetailed{}
	for _, student := range studentsList {
		if student.GraduationStatus.IsAdvisorConfirmed == 3 && student.GraduationStatus.IsDepSecConfirmed == 3 && student.GraduationStatus.IsFacultyConfirmed == 3 && student.GraduationStatus.IsStdAffConfirmed == 3 {
			filteredStudentList = append(filteredStudentList, student)
		}
	}
	return filteredStudentList
}

func FilterStudentListByAdvisor(studentsList []customized.StudentDetailed) []customized.StudentDetailed {
	filteredStudentList := []customized.StudentDetailed{}
	for _, student := range studentsList {
		if student.GraduationStatus.IsSystemConfirmed != 0 {
			filteredStudentList = append(filteredStudentList, student)
		}
	}
	return filteredStudentList
}

func FilterStudentListByDepartmentSecretary(studentsList []customized.StudentDetailed) []customized.StudentDetailed {
	filteredStudentList := []customized.StudentDetailed{}
	for _, student := range studentsList {
		if student.GraduationStatus.IsAdvisorConfirmed == 3 {
			filteredStudentList = append(filteredStudentList, student)
		}
	}
	return filteredStudentList
}

func FilterStudentListByFacultySecretary(studentsList []customized.StudentDetailed) []customized.StudentDetailed {
	filteredStudentList := []customized.StudentDetailed{}
	for _, student := range studentsList {
		if student.GraduationStatus.IsAdvisorConfirmed == 3 && student.GraduationStatus.IsDepSecConfirmed == 3 {
			filteredStudentList = append(filteredStudentList, student)
		}
	}
	return filteredStudentList
}

func FilterStudentListByStudentAffairs(studentsList []customized.StudentDetailed) []customized.StudentDetailed {
	filteredStudentList := []customized.StudentDetailed{}
	for _, student := range studentsList {
		if student.GraduationStatus.IsAdvisorConfirmed == 3 && student.GraduationStatus.IsDepSecConfirmed == 3 && student.GraduationStatus.IsFacultyConfirmed == 3 {
			filteredStudentList = append(filteredStudentList, student)
		}
	}
	return filteredStudentList
}
