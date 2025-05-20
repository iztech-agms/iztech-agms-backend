package crud

import (
	"graduation-system/globals"
	"log"
)

type Faculty struct {
	Name string `gorm:"column:name;type:varchar(255);primaryKey" json:"name"`
}

func (Faculty) TableName() string {
	return "faculties"
}

// Get all faculties
func GetFaculties() []Faculty {
	var faculties []Faculty
	if err := globals.GMSDB.Find(&faculties).Error; err != nil {
		log.Printf("(Error) : error getting faculties : %v", err)
	}
	return faculties
}

// Get faculty by name
func GetFacultyByName(name string) Faculty {
	var faculty Faculty
	if err := globals.GMSDB.Where("faculty_name = ?", name).First(&faculty).Error; err != nil {
		log.Printf("(Error) : error getting faculty : %v", err)
	}
	return faculty
}

// Create faculty
func CreateFaculty(faculty *Faculty) error {
	if err := globals.GMSDB.Create(faculty).Error; err != nil {
		log.Printf("(Error) : error creating faculty : %v", err)
		return err
	}
	return nil
}

// Update faculty
func UpdateFaculty(faculty Faculty) error {
	if err := globals.GMSDB.Save(&faculty).Error; err != nil {
		log.Printf("(Error) : error updating faculty : %v", err)
		return err
	}
	return nil
}

// Delete faculty
func DeleteFacultyByName(name string) error {
	if err := globals.GMSDB.Delete(&Faculty{}, name).Error; err != nil {
		log.Printf("(Error) : error deleting faculty : %v", err)
		return err
	}
	return nil
}
