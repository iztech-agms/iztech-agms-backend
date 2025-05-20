package crud

import (
	"graduation-system/globals"
	"log"
)

type StudentAffairs struct {
	ID   int  `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	User User `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;" json:"user"`
}

func (StudentAffairs) TableName() string {
	return "student_affairs"
}

// Get all student_affairss //Not sure what to call the plural of student affairs -Bahadır
func GetStudentAffairss() []StudentAffairs {
	var student_affairss []StudentAffairs
	if err := globals.GMSDB.Find(&student_affairss).Error; err != nil {
		log.Printf("(Error) : error getting student_affairss : %v", err)
	}
	return student_affairss
}

// Get student_affairs by ID
func GetStudentAffairsByID(id int) StudentAffairs {
	var student_affairs StudentAffairs
	if err := globals.GMSDB.Where("id = ?", id).First(&student_affairs).Error; err != nil {
		log.Printf("(Error) : error getting student_affairs : %v", err)
	}
	return student_affairs
}

// Create student_affairs
func CreateStudentAffairs(student_affairs *StudentAffairs) error {
	if err := globals.GMSDB.Create(student_affairs).Error; err != nil {
		log.Printf("(Error) : error creating student_affairs : %v", err)
		return err
	}
	return nil
}

// Update student_affairs
func UpdateStudentAffairs(student_affairs StudentAffairs) error {
	if err := globals.GMSDB.Save(&student_affairs).Error; err != nil {
		log.Printf("(Error) : error updating student_affairs : %v", err)
		return err
	}
	return nil
}

// Delete student_affairs
func DeleteStudentAffairsByID(id int) error {
	if err := globals.GMSDB.Delete(&StudentAffairs{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting student_affairs : %v", err)
		return err
	}
	return nil
}
