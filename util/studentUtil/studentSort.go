package studentUtil

import (
	"graduation-system/crud/customized"
	"slices"
)

func StudentSort(g []customized.StudentDetailed) {
	slices.SortFunc(g, func(a, b customized.StudentDetailed) int {
		// Prioritize semester ≤ 8 over semester ≥ 9
		aHighPriority := a.GraduationStatus.StudentSemester <= 8
		bHighPriority := b.GraduationStatus.StudentSemester <= 8

		// If priority differs, the high priority one comes first
		if aHighPriority != bHighPriority {
			if aHighPriority {
				return -1 // a comes before b
			}
			return 1 // b comes before a
		}

		// If semesters priorities are equal, sort by GPA descending
		if a.GraduationStatus.StudentGPA > b.GraduationStatus.StudentGPA {
			return -1
		} else if a.GraduationStatus.StudentGPA < b.GraduationStatus.StudentGPA {
			return 1
		}

		return 0
	})
}
