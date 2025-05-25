package crud

import (
	"graduation-system/globals"
	"log"
)

type Department struct {
	Name        string `gorm:"column:name;type:varchar(255);primaryKey" json:"name"`
	FacultyName string `gorm:"column:faculty_name;type:varchar(255);not null" json:"faculty_name"`
}

func (Department) TableName() string {
	return "departments"
}

// Get all departments
func GetDepartments() []Department {
	var departments []Department
	if err := globals.GMSDB.Find(&departments).Error; err != nil {
		log.Printf("(Error) : error getting departments : %v", err)
	}
	return departments
}

// Get department by name
func GetDepartmentByName(name string) Department {
	var department Department
	if err := globals.GMSDB.Where("department_name = ?", name).First(&department).Error; err != nil {
		log.Printf("(Error) : error getting department : %v", err)
	}
	return department
}

// Get department by faculty name
func GetDepartmentByFacultyName(facultyName string) []Department {
	var departments []Department
	if err := globals.GMSDB.Where("faculty_name = ?", facultyName).Find(&departments).Error; err != nil {
		log.Printf("(Error) : error getting department : %v", err)
	}
	return departments
}

// Create department
func CreateDepartment(department *Department) error {
	if err := globals.GMSDB.Create(department).Error; err != nil {
		log.Printf("(Error) : error creating department : %v", err)
		return err
	}
	return nil
}

// Update department
func UpdateDepartment(department Department) error {
	if err := globals.GMSDB.Save(&department).Error; err != nil {
		log.Printf("(Error) : error updating department : %v", err)
		return err
	}
	return nil
}

// Delete department
func DeleteDepartmentByName(name string) error {
	if err := globals.GMSDB.Delete(&Department{}, name).Error; err != nil {
		log.Printf("(Error) : error deleting department : %v", err)
		return err
	}
	return nil
}
