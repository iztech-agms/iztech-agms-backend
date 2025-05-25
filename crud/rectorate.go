package crud

import (
	"graduation-system/globals"
	"log"
)

type Rectorate struct {
	ID             int    `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	OfficeLocation string `gorm:"column:office_location;type:varchar(255);not null" json:"office_location"`
}

func (Rectorate) TableName() string {
	return "rectorate"
}

// Get all rectorates
func GetRectorates() []Rectorate {
	var rectorates []Rectorate
	if err := globals.GMSDB.Find(&rectorates).Error; err != nil {
		log.Printf("(Error) : error getting rectorates : %v", err)
	}
	return rectorates
}

// Get rectorate by ID
func GetRectorateByID(id int) Rectorate {
	var rectorate Rectorate
	if err := globals.GMSDB.Where("id = ?", id).First(&rectorate).Error; err != nil {
		log.Printf("(Error) : error getting rectorate : %v", err)
	}
	return rectorate
}

// Create rectorate
func CreateRectorate(rectorate *Rectorate) error {
	if err := globals.GMSDB.Create(rectorate).Error; err != nil {
		log.Printf("(Error) : error creating rectorate : %v", err)
		return err
	}
	return nil
}

// Update rectorate
func UpdateRectorate(rectorate Rectorate) error {
	if err := globals.GMSDB.Save(&rectorate).Error; err != nil {
		log.Printf("(Error) : error updating rectorate : %v", err)
		return err
	}
	return nil
}

// Delete rectorate
func DeleteRectorateByID(id int) error {
	if err := globals.GMSDB.Delete(&Rectorate{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting rectorate : %v", err)
		return err
	}
	return nil
}
