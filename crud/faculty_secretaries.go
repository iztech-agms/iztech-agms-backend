package crud

import (
	"graduation-system/globals"
	"log"
)

type FacultySecretary struct {
	ID          int    `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	FacultyName string `gorm:"column:faculty_name;type:varchar(255);not null" json:"faculty_name"`
	OfficeLocation string `gorm:"column:office_location;type:varchar(255);not null" json:"office_location"`
}

func (FacultySecretary) TableName() string {
	return "faculty_secretaries"
}

// Get all faculty_secretaries
func GetFacultySecretaries() []FacultySecretary {
	var faculty_secretaries []FacultySecretary
	if err := globals.GMSDB.Find(&faculty_secretaries).Error; err != nil {
		log.Printf("(Error) : error getting faculty_secretaries : %v", err)
	}
	return faculty_secretaries
}

// Get faculty_secretary by ID
func GetFacultySecretaryByID(id int) FacultySecretary {
	var faculty_secretary FacultySecretary
	if err := globals.GMSDB.Where("id = ?", id).First(&faculty_secretary).Error; err != nil {
		log.Printf("(Error) : error getting faculty_secretary : %v", err)
	}
	return faculty_secretary
}

// Create faculty_secretary
func CreateFacultySecretary(faculty_secretary *FacultySecretary) error {
	if err := globals.GMSDB.Create(faculty_secretary).Error; err != nil {
		log.Printf("(Error) : error creating faculty_secretary : %v", err)
		return err
	}
	return nil
}

// Update faculty_secretary
func UpdateFacultySecretary(faculty_secretary FacultySecretary) error {
	if err := globals.GMSDB.Save(&faculty_secretary).Error; err != nil {
		log.Printf("(Error) : error updating faculty_secretary : %v", err)
		return err
	}
	return nil
}

// Delete faculty_secretary
func DeleteFacultySecretaryByID(id int) error {
	if err := globals.GMSDB.Delete(&FacultySecretary{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting faculty_secretary : %v", err)
		return err
	}
	return nil
}
