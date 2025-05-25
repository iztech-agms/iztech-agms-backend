package customized

import (
	"graduation-system/crud"
	"graduation-system/globals"
	"log"
)

type StudentDetailed struct {
	ID               int                   `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	StudentID        int                   `gorm:"column:student_id;type:int(11);not null" json:"student_id"`
	AdvisorID        int                   `gorm:"column:advisor_id;type:int(11);not null" json:"advisor_id"`
	Advisor          crud.User             `gorm:"foreignKey:ID;references:AdvisorID" json:"advisor"`
	Student          crud.Student          `gorm:"foreignKey:StudentID;references:ID" json:"student"`
	User             crud.User             `gorm:"foreignKey:ID;references:StudentID" json:"user"`
	FacultyName      string                `gorm:"column:faculty_name;type:varchar(255);not null" json:"faculty_name"`
	DepartmentName   string                `gorm:"column:department_name;type:varchar(255);not null" json:"department_name"`
	GraduationStatus crud.GraduationStatus `gorm:"foreignKey:StudentID;references:StudentID" json:"graduation_status"`
}

func GetStudentListDetailedByUserIDs(userIDs []int) []StudentDetailed {
	studentsDetailed := []StudentDetailed{}
	if err := globals.GMSDB.Model(&crud.Student{}).
		Select("students.*, graduation_statuses.*, users.*, departments.name as department_name, departments.faculty_name as faculty_name").
		Joins("JOIN graduation_statuses ON students.id = graduation_statuses.student_id").
		Joins("JOIN users ON students.id = users.id").
		Joins("JOIN advisors ON students.advisor_id = advisors.id").
		Joins("JOIN departments ON advisors.department_name = departments.name").
		Where("students.id IN (?)", userIDs).
		Order("graduation_statuses.student_gpa desc").
		Preload("Student").
		Preload("User").
		Preload("GraduationStatus").
		Preload("Advisor").
		Find(&studentsDetailed).Error; err != nil {
		log.Printf("(Error) : error getting student list detailed by user ids : %v", err)
	}
	return studentsDetailed
}

// Get student IDs by user ID
func GetStudentIDsByUserID(userID int) []int {
	studentIDs := []int{}

	// First get user
	user := crud.GetUserByID(userID)
	if user.ID == 0 {
		log.Printf("User not found with id: %d", userID)
		return studentIDs
	}

	// If role == 'advisor'
	if user.Role == "advisor" {
		if err := globals.GMSDB.Model(&crud.Student{}).
			Where("advisor_id = ?", userID).
			Pluck("id", &studentIDs).Error; err != nil {
			log.Printf("(Error) : error getting student ids by advisor id : %v", err)
		}
		return studentIDs
	}

	// If role == 'department_secretary' || 'faculty_secretary' || 'student_affairs'
	if user.Role == "department_secretary" {
		// One query to get all students in the department
		if err := globals.GMSDB.Model(&crud.Student{}).
			Joins("JOIN advisors ON students.advisor_id = advisors.id").
			Joins("JOIN department_secretaries ON department_secretaries.department_name = advisors.department_name").
			Where("department_secretaries.id = ?", userID).
			Pluck("students.id", &studentIDs).Error; err != nil {
			log.Printf("(Error) : error getting students by department secretary : %v", err)
		}
		return studentIDs
	}

	if user.Role == "faculty_secretary" {
		// One query to get all students in the faculty
		if err := globals.GMSDB.Model(&crud.Student{}).
			Joins("JOIN advisors ON students.advisor_id = advisors.id").
			Joins("JOIN departments ON departments.name = advisors.department_name").
			Joins("JOIN faculty_secretaries ON faculty_secretaries.faculty_name = departments.faculty_name").
			Where("faculty_secretaries.id = ?", userID).
			Pluck("students.id", &studentIDs).Error; err != nil {
			log.Printf("(Error) : error getting students by faculty secretary : %v", err)
		}
		return studentIDs
	}

	if user.Role == "student_affairs" {
		// One query to get all students
		if err := globals.GMSDB.Model(&crud.Student{}).
			Pluck("id", &studentIDs).Error; err != nil {
			log.Printf("(Error) : error getting all students : %v", err)
		}
		return studentIDs
	}

	return studentIDs
}
