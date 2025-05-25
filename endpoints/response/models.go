package response

import (
	"graduation-system/crud"
	"graduation-system/crud/customized"
)

type ResponseMessage struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type StudentsResp struct {
	Status   ResponseMessage `json:"status"`
	Students []crud.User     `json:"students"`
}

type StudentsDetailedResp struct {
	Status   ResponseMessage              `json:"status"`
	Students []customized.StudentDetailed `json:"students"`
}

type NotificationsResp struct {
	Status        ResponseMessage     `json:"status"`
	Notifications []crud.Notification `json:"notifications"`
}

type GraduationStatusResp struct {
	Status           ResponseMessage         `json:"status"`
	GraduationStatus []crud.GraduationStatus `json:"graduation_status"`
}

type GraduationYearResp struct {
	Status   ResponseMessage       `json:"status"`
	GradYear []crud.GraduationYear `json:"graduation_years"`
}
