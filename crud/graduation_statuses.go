package crud

import (
	"graduation-system/globals"
	"log"
)

type GraduationStatus struct {
	ID                 int  `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	Year               int  `gorm:"column:id;type:int(13);not null" json:"year"`
	StudentID          int  `gorm:"column:student_id;type:int(11);not null" json:"student_id"`
	StudentSemester    int  `gorm:"column:student_semester;type:int(5);not null" json:"student_semester"`
	StudentGPA         float64  `gorm:"column:student_gpa;type:double;not null" json:"student_gpa"` //TODO change to float
	IsAdvisorConfirmed bool `gorm:"not null" json:"is_advisor_confirmed"`
	IsDepSecConfirmed  bool `gorm:"not null" json:"is_dep_sec_confirmed"`
	IsFacultyConfirmed bool `gorm:"not null" json:"is_faculty_confirmed"`
	IsStdAffConfirmed  bool `gorm:"not null" json:"is_std_aff_confirmed"`
}

func (GraduationStatus) TableName() string {
	return "graduation_statuses"
}

// Get all graduation_statuses
func GetGraduationStatuses() []GraduationStatus {
	var graduation_statuses []GraduationStatus
	if err := globals.GMSDB.Find(&graduation_statuses).Error; err != nil {
		log.Printf("(Error) : error getting graduation_statuses : %v", err)
	}
	return graduation_statuses
}

// Get grad statuses by year
func GetGraduationStatusesByYear(year int) []GraduationStatus {
	var graduation_statuses []GraduationStatus
	if err := globals.GMSDB.Find(&graduation_statuses).Where("year = ?", year).Error; err != nil {
		log.Printf("(Error) : error getting graduation_statuses : %v", err)
	}
	return graduation_statuses
}

// Get graduation_status by ID
func GetGraduationStatusByID(id int) GraduationStatus {
	var graduation_status GraduationStatus
	if err := globals.GMSDB.Where("id = ?", id).First(&graduation_status).Error; err != nil {
		log.Printf("(Error) : error getting graduation_status : %v", err)
	}
	return graduation_status
}

// Get gradStatus by matching student ID
func GetGraduationStatusByStudentID(studentID int) GraduationStatus {
	var res GraduationStatus
	if err := globals.GMSDB.Where("student_id = ?", studentID).First(&res).Error; err != nil {
		log.Printf("(Error) : error getting graduation_status : %v", err)
	}
	return res
}


// Create graduation_status
func CreateGraduationStatus(graduation_status *GraduationStatus) error {
	if err := globals.GMSDB.Create(graduation_status).Error; err != nil {
		log.Printf("(Error) : error creating graduation_status : %v", err)
		return err
	}
	return nil
}

// Update graduation_status
func UpdateGraduationStatus(graduation_status GraduationStatus) error {
	if err := globals.GMSDB.Save(&graduation_status).Error; err != nil {
		log.Printf("(Error) : error updating graduation_status : %v", err)
		return err
	}
	return nil
}

// Delete graduation_status
func DeleteGraduationStatusByID(id int) error {
	if err := globals.GMSDB.Delete(&GraduationStatus{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting graduation_status : %v", err)
		return err
	}
	return nil
}
