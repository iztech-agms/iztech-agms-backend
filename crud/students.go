package crud

import (
	"graduation-system/globals"
	"log"
)

type Student struct {
	ID        int `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	AdvisorID int `gorm:"column:advisor_id;type:int(11);not null" json:"advisor_id"`
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

// Get student by ID
func GetStudentByID(id int) Student {
	var student Student
	if err := globals.GMSDB.Where("id = ?", id).First(&student).Error; err != nil {
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
