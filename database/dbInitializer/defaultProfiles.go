package dbinitializer

import (
	"encoding/json"
	"graduation-system/crud"
	"graduation-system/util"
	"io"
	"log"
	"os"
)

type GraduationStatusJSON struct {
	Year               int     `json:"year"`
	StudentUsername    string  `json:"student_username"` // Username instead of id
	StudentSemester    int     `json:"student_semester"`
	StudentGPA         float64 `json:"student_gpa"`
	StudentECTS        int     `json:"student_ects"`
	IsSystemConfirmed  int     `json:"is_system_confirmed"`
	IsAdvisorConfirmed int     `json:"is_advisor_confirmed"`
	IsDepSecConfirmed  int     `json:"is_dep_sec_confirmed"`
	IsFacultyConfirmed int     `json:"is_faculty_confirmed"`
	IsStdAffConfirmed  int     `json:"is_std_aff_confirmed"`
}

type UserJSON struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Role            string `json:"role"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Telephone       string `json:"telephone"`
	AdvisorUsername string `json:"advisor_username"`
	DepartmentName  string `json:"department_name"`
	FacultyName     string `json:"faculty_name"`
	OfficeLocation  string `json:"office_location"`
	EnrollmentYear  int    `json:"enrollment_year"`
}

func importUsersFromJSON(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var users []UserJSON
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return err
	}

	for _, u := range users {
		hashed_password, err := util.HashPassword(u.Password)
		if err != nil {
			log.Printf("Hashing password failed during default user creation: %v", err)
			continue
		}
		user := &crud.User{
			Username:  u.Username,
			Email:     u.Email,
			Password:  hashed_password,
			Telephone: u.Telephone,
			Role:      u.Role,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}

		crud.CreateUser(user)

		switch u.Role {
		case "student":
			advisor := crud.GetAdvisorByUsername(u.AdvisorUsername)
			crud.CreateStudent(&crud.Student{
				ID:             user.ID,
				AdvisorID:      advisor.ID,
				EnrollmentYear: u.EnrollmentYear,
			})
		case "advisor":
			crud.CreateAdvisor(&crud.Advisor{
				ID:             user.ID,
				DepartmentName: u.DepartmentName,
				OfficeLocation: u.OfficeLocation,
			})
		case "department_secretary":
			crud.CreateDepartmentSecretary(&crud.DepartmentSecretary{
				ID:             user.ID,
				DepartmentName: u.DepartmentName,
				OfficeLocation: u.OfficeLocation,
			})
		case "faculty_secretary":
			crud.CreateFacultySecretary(&crud.FacultySecretary{
				ID:             user.ID,
				FacultyName:    u.FacultyName,
				OfficeLocation: u.OfficeLocation,
			})
		case "student_affairs":
			crud.CreateStudentAffairs(&crud.StudentAffairs{
				ID:             user.ID,
				OfficeLocation: u.OfficeLocation,
			})
		case "rectorate":
			crud.CreateRectorate(&crud.Rectorate{
				ID:             user.ID,
				OfficeLocation: u.OfficeLocation,
			})
		}
	}

	return nil
}

func importGraduationStatusesFromJSON(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var statuses []GraduationStatusJSON
	if err := json.Unmarshal(bytes, &statuses); err != nil {
		return err
	}

	for _, s := range statuses {
		student := crud.GetStudentByUsername(s.StudentUsername)
		if student.ID == 0 {
			continue
		}

		status := &crud.GraduationStatus{
			Year:               s.Year,
			StudentID:          student.ID,
			StudentSemester:    s.StudentSemester,
			StudentGPA:         s.StudentGPA,
			StudentECTS:        s.StudentECTS,
			IsSystemConfirmed:  s.IsSystemConfirmed,
			IsAdvisorConfirmed: s.IsAdvisorConfirmed,
			IsDepSecConfirmed:  s.IsDepSecConfirmed,
			IsFacultyConfirmed: s.IsFacultyConfirmed,
			IsStdAffConfirmed:  s.IsStdAffConfirmed,
		}
		crud.CreateGraduationStatus(status)
	}

	return nil
}

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
		crud.CreateRole(&crud.Role{
			Name: "rectorate",
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

	if len(crud.GetUsers()) == 0 {
		advisor_path := "../database/dbInitializer/defaultUsers/advisors.json"
		department_secretary_path := "../database/dbInitializer/defaultUsers/department_secretaries.json"
		faculty_secretary_path := "../database/dbInitializer/defaultUsers/faculty_secretaries.json"
		student_affairs_path := "../database/dbInitializer/defaultUsers/student_affairs.json"
		rectorate_path := "../database/dbInitializer/defaultUsers/rectorates.json"
		students_path := "../database/dbInitializer/defaultUsers/students.json"
		graduation_status_path := "../database/dbInitializer/defaultUsers/graduation_statuses.json"
		err := importUsersFromJSON(advisor_path)
		if err != nil {
			log.Printf("Error users from file %s: %v\n", advisor_path, err)
		}
		err = importUsersFromJSON(department_secretary_path)
		if err != nil {
			log.Printf("Error users from file %s: %v\n", department_secretary_path, err)
		}
		err = importUsersFromJSON(faculty_secretary_path)
		if err != nil {
			log.Printf("Error users from file %s: %v\n", faculty_secretary_path, err)
		}
		err = importUsersFromJSON(student_affairs_path)
		if err != nil {
			log.Printf("Error users from file %s: %v\n", student_affairs_path, err)
		}
		err = importUsersFromJSON(rectorate_path)
		if err != nil {
			log.Printf("Error users from file %s: %v\n", rectorate_path, err)
		}
		err = importUsersFromJSON(students_path)
		if err != nil {
			log.Printf("Error users from file %s: %v\n", students_path, err)
		}

		err = importGraduationStatusesFromJSON(graduation_status_path)
		if err != nil {
			log.Printf("Error loading graduation statuses from file %s: %v\n", graduation_status_path, err)
		}
	}

}
