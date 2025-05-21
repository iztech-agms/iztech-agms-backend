package crud

import (
	"graduation-system/globals"
	"log"
	"time"
)

type Notification struct {
	ID                 int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	UserID             int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	CreatedAt          time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	IsNotificationRead bool      `gorm:"column:is_notification_read;type:tinyint(1);not null" json:"is_notification_read"`
	Message            string    `gorm:"column:message;type:text;not null" json:"message"`
	Title              string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
}

func (Notification) TableName() string {
	return "notifications"
}

// Get all notifications
func GetNotifications() []Notification {
	var notifications []Notification
	if err := globals.GMSDB.Find(&notifications).Error; err != nil {
		log.Printf("(Error) : error getting notifications : %v", err)
	}
	return notifications
}

// Get notification by ID
func GetNotificationByID(id int) Notification {
	var notification Notification
	if err := globals.GMSDB.Where("id = ?", id).First(&notification).Error; err != nil {
		log.Printf("(Error) : error getting notification : %v", err)
	}
	return notification
}

// Get notifications by reciever ID
func GetNotificationsByRecieverID(recieverID int) []Notification {
	notifications := []Notification{}
	if err := globals.GMSDB.Where("user_id = ?", recieverID).Find(&notifications).Error; err != nil {
		log.Printf("(Error) : error getting notifications : %v", err)
	}
	return notifications
}

// Create notification
func CreateNotification(notification *Notification) error {
	if err := globals.GMSDB.Create(notification).Error; err != nil {
		log.Printf("(Error) : error creating notification : %v", err)
		return err
	}
	return nil
}

// Update notification
func UpdateNotification(notification Notification) error {
	if err := globals.GMSDB.Save(&notification).Error; err != nil {
		log.Printf("(Error) : error updating notification : %v", err)
		return err
	}
	return nil
}

// Delete notification
func DeleteNotificationByID(id int) error {
	if err := globals.GMSDB.Delete(&Notification{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting notification : %v", err)
		return err
	}
	return nil
}
