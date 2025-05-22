package dbinitializer

import (
	"graduation-system/crud"
)

func InitializeDefaultProfiles() {
	if len(crud.GetRoles()) == 0 {
		// Default Roles
		crud.CreateRole(&crud.Role{
			Name: "student",
		})
		crud.CreateRole(&crud.Role{
			Name: "advisor",
		})
		crud.CreateRole(&crud.Role{
			Name: "department_secretary",
		})
		crud.CreateRole(&crud.Role{
			Name: "faculty_secretary",
		})
		crud.CreateRole(&crud.Role{
			Name: "student_affairs",
		})
	}

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
		user := &crud.User{
			Username:  "stdaff1",
			Password:  "1",
			Telephone: "todo",
			Role:      "student_affairs",
			FirstName: "William",
			LastName:  "Johnson",
		}
		crud.CreateUser(user)
		crud.CreateStudentAffairs(&crud.StudentAffairs{
			ID:             user.ID,
			OfficeLocation: "todo",
		})
	}

	if len(crud.GetFacultySecretaries()) == 0 {
		// Default faculty secretaries
		user := &crud.User{
			Username:  "facsec1",
			Password:  "1",
			Telephone: "todo",
			Role:      "faculty_secretary",
			FirstName: "Robert",
			LastName:  "Johnson",
		}
		crud.CreateUser(user)
		crud.CreateFacultySecretary(&crud.FacultySecretary{
			ID:             user.ID,
			FacultyName:    "Engineering",
			OfficeLocation: "todo",
		})
	}

	if len(crud.GetDepartmentSecretaries()) == 0 {
		// Default department secretaries
		user := &crud.User{
			Username:  "depsec1",
			Password:  "1",
			Telephone: "todo",
			Role:      "department_secretary",
			FirstName: "Bob",
			LastName:  "Johnson",
		}
		crud.CreateUser(user)
		crud.CreateDepartmentSecretary(&crud.DepartmentSecretary{
			ID:             user.ID,
			DepartmentName: "Computer Engineering",
			OfficeLocation: "todo",
		})
	}

	if len(crud.GetAdvisors()) == 0 {

		// Default advisors
		user := &crud.User{
			Username:  "buketoksuzoglu@iyte.edu.tr",
			Password:  "4321",
			Telephone: "+90 232 750 7864",
			Role:      "advisor",
			FirstName: "Buket",
			LastName:  "Erşahin",
		}

		crud.CreateUser(user)
		// Now use the auto-generated user.ID for the advisor
		crud.CreateAdvisor(&crud.Advisor{
			ID:             user.ID,
			DepartmentName: "Computer Engineering",
			OfficeLocation: "todo",
		})

		user = &crud.User{
			Username:  "290201064",
			Password:  "samet123",
			Telephone: "todo",
			Role:      "student",
			FirstName: "Samet",
			LastName:  "Hodaman",
		}
		crud.CreateUser(user)
		crud.CreateStudent(&crud.Student{
			ID:        user.ID,
			AdvisorID: crud.GetAdvisorByUsername("buketoksuzoglu@iyte.edu.tr").ID,
		})

		crud.CreateNotification(&crud.Notification{
			UserID:             user.ID,
			IsNotificationRead: false,
			Title:              "Welcome to the system",
			Message:            "Enjoy the system!	",
		})

		user = &crud.User{
			Username:  "300201079",
			Password:  "1234",
			Telephone: "todo",
			Role:      "student",
			FirstName: "Bahadir Efe",
			LastName:  "AVSAR",
		}
		crud.CreateUser(user)
		crud.CreateStudent(&crud.Student{
			ID:        user.ID,
			AdvisorID: crud.GetAdvisorByUsername("buketoksuzoglu@iyte.edu.tr").ID,
		})
	}
	if len(crud.GetGraduationStatuses()) == 0 {
		crud.CreateGraduationStatus(&crud.GraduationStatus{
			Year:               2025,
			StudentID:          crud.GetStudentByUsername("290201064").ID,
			StudentSemester:    8,
			StudentGPA:         2.5,
			StudentCredits:     121,
			IsAdvisorConfirmed: 2,
			IsDepSecConfirmed:  2,
			IsFacultyConfirmed: 2,
			IsStdAffConfirmed:  2,
		})
		crud.CreateGraduationStatus(&crud.GraduationStatus{
			Year:               2025,
			StudentID:          crud.GetStudentByUsername("300201079").ID,
			StudentSemester:    8,
			StudentGPA:         2.6,
			StudentCredits:     120,
			IsAdvisorConfirmed: 2,
			IsDepSecConfirmed:  2,
			IsFacultyConfirmed: 2,
			IsStdAffConfirmed:  2,
		})
	}

}
