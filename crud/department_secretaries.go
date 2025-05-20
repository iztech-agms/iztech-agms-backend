package crud

import (
	"graduation-system/globals"
	"log"
)

type DepartmentSecretary struct {
	ID             int    `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	DepartmentName string `gorm:"column:department_name;type:varchar(255);not null" json:"department_name"`

	User       User       `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;" json:"user"`
	Department Department `gorm:"foreignKey:DepartmentName;references:Name" json:"department"`
}

func (DepartmentSecretary) TableName() string {
	return "department_secretaries"
}

// Get all department_secretaries
func GetDepartmentSecretaries() []DepartmentSecretary {
	var department_secretaries []DepartmentSecretary
	if err := globals.GMSDB.Find(&department_secretaries).Error; err != nil {
		log.Printf("(Error) : error getting department_secretaries : %v", err)
	}
	return department_secretaries
}

// Get department_secretary by ID
func GetDepartmentSecretaryByID(id int) DepartmentSecretary {
	var department_secretary DepartmentSecretary
	if err := globals.GMSDB.Where("id = ?", id).First(&department_secretary).Error; err != nil {
		log.Printf("(Error) : error getting department_secretary : %v", err)
	}
	return department_secretary
}

// Create department_secretary
func CreateDepartmentSecretary(department_secretary *DepartmentSecretary) error {
	if err := globals.GMSDB.Create(department_secretary).Error; err != nil {
		log.Printf("(Error) : error creating department_secretary : %v", err)
		return err
	}
	return nil
}

// Update department_secretary
func UpdateDepartmentSecretary(department_secretary DepartmentSecretary) error {
	if err := globals.GMSDB.Save(&department_secretary).Error; err != nil {
		log.Printf("(Error) : error updating department_secretary : %v", err)
		return err
	}
	return nil
}

// Delete department_secretary
func DeleteDepartmentSecretaryByID(id int) error {
	if err := globals.GMSDB.Delete(&DepartmentSecretary{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting department_secretary : %v", err)
		return err
	}
	return nil
}
