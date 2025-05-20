package response

import "graduation-system/crud"

type ResponseMessage struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type StudentsResp struct {
	Status   ResponseMessage `json:"status"`
	Students []crud.User     `json:"students"`
}

type NotificationsResp struct {
	Status        ResponseMessage     `json:"status"`
	Notifications []crud.Notification `json:"notifications"`
}
