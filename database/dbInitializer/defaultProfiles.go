package dbinitializer

import (
	"graduation-system/crud"
)

func InitializeDefaultProfiles() {

	if len(crud.GetFaculties()) == 0 {
		// Default faculties
		crud.CreateFaculty(&crud.Faculty{
			Name: "Engineering",
		})

		crud.CreateFaculty(&crud.Faculty{
			Name: "Science",
		})

		crud.CreateFaculty(&crud.Faculty{
			Name: "Architecture",
		})
	}
	if len(crud.GetDepartments()) == 0 {
		// Default departments found in https://en.iyte.edu.tr/academic/academic-units/
		//Engineering
		crud.CreateDepartment(&crud.Department{
			Name:        "Computer Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Bioengineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Environmental Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Energy Systems Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Electrical-Electronics Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Food Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Civil Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Chemical Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Mechanical Engineering",
			FacultyName: "Engineering",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Materials Science and Engineering",
			FacultyName: "Engineering",
		})
		//Science
		crud.CreateDepartment(&crud.Department{
			Name:        "Physics",
			FacultyName: "Science",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Photonics",
			FacultyName: "Science",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Chemistry",
			FacultyName: "Science",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Mathematics",
			FacultyName: "Science",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Molecular Biology and Genetics",
			FacultyName: "Science",
		})
		//Architecture
		crud.CreateDepartment(&crud.Department{
			Name:        "Industrial Design",
			FacultyName: "Architecture",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Conservation and Restoration Cultural Heritage",
			FacultyName: "Architecture",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "Architecture",
			FacultyName: "Architecture",
		})
		crud.CreateDepartment(&crud.Department{
			Name:        "City and Regional Planning",
			FacultyName: "Architecture",
		})
	}

	if len(crud.GetStudentAffairss()) == 0 {
		// Default Student affairs
		crud.CreateUser(&crud.User{
			ID:        5,
			Username:  "stdaff1",
			Password:  "1",
			Role:      "student_affairs",
			FirstName: "William",
			LastName:  "Johnson",
		})
		crud.CreateStudentAffairs(&crud.StudentAffairs{
			ID: 5,
		})
	}

	if len(crud.GetFacultySecretaries()) == 0 {
		// Default faculty secretaries
		crud.CreateUser(&crud.User{
			ID:        4,
			Username:  "facsec1",
			Password:  "1",
			Role:      "faculty_secretary",
			FirstName: "Robert",
			LastName:  "Johnson",
		})
		crud.CreateFacultySecretary(&crud.FacultySecretary{
			ID:          4,
			FacultyName: "Engineering",
		})
	}

	if len(crud.GetDepartmentSecretaries()) == 0 {
		// Default department secretaries
		crud.CreateUser(&crud.User{
			ID:        3,
			Username:  "depsec1",
			Password:  "1",
			Role:      "department_secretary",
			FirstName: "Bob",
			LastName:  "Johnson",
		})
		crud.CreateDepartmentSecretary(&crud.DepartmentSecretary{
			ID:             3,
			DepartmentName: "Computer Engineering",
		})
	}

	if len(crud.GetAdvisors()) == 0 {

		// Default advisors
		crud.CreateUser(&crud.User{
			ID:        1,
			Username:  "advisor1",
			Password:  "4321",
			Role:      "advisor",
			FirstName: "John",
			LastName:  "Johnson",
		})
		crud.CreateAdvisor(&crud.Advisor{
			ID:             1,
			DepartmentName: "Computer Engineering",
		})

		crud.CreateUser(&crud.User{
			ID:        2,
			Username:  "290201064",
			Password:  "samet123",
			Role:      "student",
			FirstName: "Samet",
			LastName:  "Hodaman",
		})
		crud.CreateStudent(&crud.Student{
			ID:        2,
			AdvisorID: 1,
		})

		crud.CreateNotification(&crud.Notification{
			UserID:             2,
			IsNotificationRead: false,
			Title:              "Welcome to the system",
			Message:            "Enjoy the system!	",
		})
	}
	if len(crud.GetGraduationStatuses()) == 0 {
		crud.CreateGraduationStatus(&crud.GraduationStatus{
			ID:                 1,
			Year:               2025,
			StudentID:          2,
			StudentSemester:    8,
			StudentGPA:         2.5,
			IsAdvisorConfirmed: false,
			IsDepSecConfirmed:  false,
			IsFacultyConfirmed: false,
			IsStdAffConfirmed:  false,
		})
	}

}
