package crud

import (
	"graduation-system/globals"
	"log"
	"time"
)

type GraduationYear struct {
	Year          int       `gorm:"column:year;type:int(13);not null;primaryKey" json:"year"`
	StartDate     time.Time `gorm:"column:start_date;type:timestamp;not null" json:"start_date"`
	EndDate       time.Time `gorm:"column:end_date;type:timestamp;not null" json:"end_date"`
	GraduateCount int       `gorm:"column:graduate_count;type:int(13);not null" json:"graduate_count"`
}

func (GraduationYear) TableName() string {
	return "graduation_years"
}

// Get all graduation_years
func GetGraduationYears() []GraduationYear {
	var graduation_years []GraduationYear
	if err := globals.GMSDB.Find(&graduation_years).Error; err != nil {
		log.Printf("(Error) : error getting graduation_years : %v", err)
	}
	return graduation_years
}

// Get graduation_year by ID
func GetGraduationYearByYear(year int) GraduationYear {
	var graduation_year GraduationYear
	if err := globals.GMSDB.Where("year = ?", year).First(&graduation_year).Error; err != nil {
		//log.Printf("(Error) : error getting graduation_year : %v", err)
	}
	return graduation_year
}

// Create graduation_year
func CreateGraduationYear(graduation_year *GraduationYear) error {
	if err := globals.GMSDB.Create(graduation_year).Error; err != nil {
		log.Printf("(Error) : error creating graduation_year : %v", err)
		return err
	}
	return nil
}

// Update graduation_year
func UpdateGraduationYear(graduation_year GraduationYear) error {
	if err := globals.GMSDB.Save(&graduation_year).Error; err != nil {
		log.Printf("(Error) : error updating graduation_year : %v", err)
		return err
	}
	return nil
}

// Delete graduation_year
func DeleteGraduationYearByID(id int) error {
	if err := globals.GMSDB.Delete(&GraduationYear{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting graduation_year : %v", err)
		return err
	}
	return nil
}
