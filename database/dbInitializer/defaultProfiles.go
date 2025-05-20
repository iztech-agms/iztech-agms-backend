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

	var user crud.User
	if len(crud.GetAdvisors()) == 0 {
		// Default advisors

		user = crud.User{
			Username:  "advisor1",
			Password:  "4321",
			Role:      "advisor",
			FirstName: "John",
			LastName:  "Johnson",
		}

		crud.CreateUser(&user)
		crud.CreateAdvisor(&crud.Advisor{
			DepartmentName: "Computer Engineering",
			User: user,
		})

		user = crud.User{
			Username:  "290201064",
			Password:  "samet123",
			Role:      "student",
			FirstName: "Samet",
			LastName:  "Hodaman",
		}
		crud.CreateUser(&user)
		crud.CreateStudent(&crud.Student{
			AdvisorID: crud.GetAdvisorByUsername("advisor1").ID,
			User: user,
		})
	}

}
