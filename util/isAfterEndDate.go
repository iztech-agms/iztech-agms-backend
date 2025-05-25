package util

import (
	"graduation-system/crud"
	"time"
)

func IsAfterEndDate() bool {
	grad_year := crud.GetGraduationYearByYear(time.Now().Year())
	if grad_year.Year == 0 {
		return true
	}
	return grad_year.EndDate.After(time.Now())
}
