package crud

import (
	"graduation-system/globals"
	"log"
)

type User struct {
	ID        int    `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	FirstName string `gorm:"column:first_name;type:varchar(255);not null" json:"first_name"`
	LastName  string `gorm:"column:last_name;type:varchar(255);not null" json:"last_name"`
	Username  string `gorm:"column:username;type:varchar(255);not null;unique" json:"username"`
	Password  string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Role      string `gorm:"column:role;type:varchar(255);not null" json:"role"` //Not sure what to do with this -Bahadır
}

func (User) TableName() string {
	return "users"
}

// Get all users
func GetUsers() []User {
	var users []User
	if err := globals.GMSDB.Find(&users).Error; err != nil {
		log.Printf("(Error) : error getting users : %v", err)
	}
	return users
}

// Get user by ID
func GetUserByID(id int) User {
	var user User
	if err := globals.GMSDB.Where("id = ?", id).First(&user).Error; err != nil {
		log.Printf("(Error) : error getting user : %v", err)
	}
	return user
}

// Get user by Username
func GetUserByUsername(username string) User {
	var user User
	if err := globals.GMSDB.Where("username = ?", username).First(&user).Error; err != nil {
		log.Printf("(Error) : error getting user : %v", err)
	}
	return user
}

// Create user
func CreateUser(user *User) error {
	if err := globals.GMSDB.Create(user).Error; err != nil {
		log.Printf("(Error) : error creating user : %v", err)
		return err
	}
	return nil
}

// Update user
func UpdateUser(user User) error {
	if err := globals.GMSDB.Save(&user).Error; err != nil {
		log.Printf("(Error) : error updating user : %v", err)
		return err
	}
	return nil
}

// Delete user
func DeleteUserByID(id int) error {
	if err := globals.GMSDB.Delete(&User{}, id).Error; err != nil {
		log.Printf("(Error) : error deleting user : %v", err)
		return err
	}
	return nil
}
