package crud

import (
	"graduation-system/globals"
	"log"
)

type Student struct {
	ID             int `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	AdvisorID      int `gorm:"column:advisor_id;type:int(11);not null" json:"advisor_id"`
	EnrollmentYear int `gorm:"column:enrollment_year;type:int(13);not null" json:"enrollment_year"`
}

func (Student) TableName() string {
	return "students"
}

// Get all students
func GetStudents() []Student {
	var students []Student
	if err := globals.GMSDB.Find(&students).Error; err != nil {
		log.Printf("(Error) : error getting students : %v", err)
	}
	return students
}

func GetStudentsByAdvisorID(advisor_id int) []Student {
	var students []Student
	if err := globals.GMSDB.Find(&students, "advisor_id = ?", advisor_id).Error; err != nil {
		log.Printf("(Error) : error getting students : %v", err)
	}
	return students //advisor department secratery
}

func GetStudentsByDepartmentSecretaryID(sec_id int) []Student {
	var students []Student

	err := globals.GMSDB.Raw(`
		SELECT students.*
		FROM department_secretaries
		JOIN advisors ON advisors.department_name = department_secretaries.department_name
		JOIN students ON students.advisor_id = advisors.id
		WHERE department_secretaries.id = ?`, sec_id).Scan(&students).Error
	if err != nil {
		log.Printf("(Error) : error getting students : %v", err)
	}
	return students //advisor department secratery
}

func GetStudentIDsByDepartmentName(name string) []int {
	var student_ids []int

	err := globals.GMSDB.Raw(`
		SELECT students.id
		FROM departments
		JOIN advisors ON advisors.department_name = departments.name
		JOIN students ON students.advisor_id = advisors.id
		WHERE departments.name= ?`, name).Scan(&student_ids).Error
	if err != nil {
		log.Printf("(Error) : error getting students : %v", err)
	}
	return student_ids
}

func GetStudentIDs() []int {
	var student_ids []int

	err := globals.GMSDB.Raw(`
		SELECT id
		FROM students`).Scan(&student_ids).Error
	if err != nil {
		log.Printf("(Error) : error getting students : %v", err)
	}
	return student_ids
}

func GetStudentsByFacultySecretaryID(sec_id int) []Student {
	var students []Student

	err := globals.GMSDB.Raw(`
		SELECT students.*
		FROM faculty_secretaries
		JOIN faculties ON faculties.faculty_name = faculty_secretaries.faculty_name
		JOIN departments ON departments.faculty_name = faculties.faculty_name
		JOIN advisors ON advisors.department_name = departments.department_name
		JOIN students ON students.advisor_id = advisors.id
		WHERE faculty_secretaries.id = ?
	`, sec_id).Scan(&students).Error

	if err != nil {
		log.Printf("Error getting students %v", err)
	}
	return students
}

// Get student by ID
func GetStudentByID(id int) Student {
	var student Student
	if err := globals.GMSDB.Where("id = ?", id).First(&student).Error; err != nil {
		log.Printf("(Error) : error getting student : %v", err)
	}
	return student
}

func GetStudentByUsername(username string) Student {
	var student Student
	err := globals.GMSDB.Table("students").
		Joins("JOIN users ON users.id = students.id").
		Where("users.username = ?", username).
		First(&student).Error

	if err != nil {
		log.Printf("(Error) : error getting student : %v", err)
	}
	return student
}

// Create student
func CreateStudent(student *Student) error {
	if err := globals.GMSDB.Create(student).Error; err != nil {
		log.Printf("(Error) : error creating student : %v", err)
		return err
	}
	return nil
}

// Update student
func UpdateStudent(student Student) error {
	if err := globals.GMSDB.Save(&student).Error; err != nil {
		log.Printf("(Error) : error updating student : %v", err)
		return err
	}
	return nil
}

// Delete student
func DeleteStudentByID(id int) error {
	if err := globals.GMSDB.Delete(&Student{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting student : %v", err)
		return err
	}
	return nil
}
