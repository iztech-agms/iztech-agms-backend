package crud

import (
	"graduation-system/globals"
	"log"
)

type Advisor struct {
	ID             int    `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	DepartmentName string `gorm:"column:department_name;type:varchar(255);not null" json:"department_name"`

	User       User       `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;" json:"user"`
	Department Department `gorm:"foreignKey:DepartmentName;references:Name" json:"department"`
}

func (Advisor) TableName() string {
	return "advisors"
}

// Get all advisors
func GetAdvisors() []Advisor {
	var advisors []Advisor
	if err := globals.GMSDB.Find(&advisors).Error; err != nil {
		log.Printf("(Error) : error getting advisors : %v", err)
	}
	return advisors
}

// Get advisor by ID
func GetAdvisorByID(id int) Advisor {
	var advisor Advisor
	if err := globals.GMSDB.Where("id = ?", id).First(&advisor).Error; err != nil {
		log.Printf("(Error) : error getting advisor : %v", err)
	}
	return advisor
}

// Get advisor by name (not tested yet)

func GetAdvisorByUsername(username string) Advisor {
	var advisor Advisor
	err := globals.GMSDB.Table("advisors").
		Joins("JOIN users ON users.id = advisors.id").
		Where("users.username = ?", username).
		First(&advisor).Error

	if err != nil {
		log.Printf("(Error) : error getting advisor : %v", err)
	}
	return advisor
}

// Create advisor
func CreateAdvisor(advisor *Advisor) error {
	if err := globals.GMSDB.Create(advisor).Error; err != nil {
		log.Printf("(Error) : error creating advisor : %v", err)
		return err
	}
	return nil
}

// Update advisor
func UpdateAdvisor(advisor Advisor) error {
	if err := globals.GMSDB.Save(&advisor).Error; err != nil {
		log.Printf("(Error) : error updating advisor : %v", err)
		return err
	}
	return nil
}

// Delete advisor
func DeleteAdvisorByID(id int) error {
	if err := globals.GMSDB.Delete(&Advisor{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting advisor : %v", err)
		return err
	}
	return nil
}
