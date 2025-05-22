package crud

import (
	"graduation-system/globals"
	"log"
)

type Role struct {
	Name string `gorm:"column:name;type:varchar(255);primaryKey" json:"name"`
}

func (Role) TableName() string {
	return "roles"
}

// Get all roles
func GetRoles() []Role {
	var roles []Role
	if err := globals.GMSDB.Find(&roles).Error; err != nil {
		log.Printf("(Error) : error getting roles : %v", err)
	}
	return roles
}

// Get role by name
func GetRoleByName(name string) Role {
	var role Role
	if err := globals.GMSDB.Where("role_name = ?", name).First(&role).Error; err != nil {
		log.Printf("(Error) : error getting role : %v", err)
	}
	return role
}

// Create role
func CreateRole(role *Role) error {
	if err := globals.GMSDB.Create(role).Error; err != nil {
		log.Printf("(Error) : error creating role : %v", err)
		return err
	}
	return nil
}

// Update role
func UpdateRole(role Role) error {
	if err := globals.GMSDB.Save(&role).Error; err != nil {
		log.Printf("(Error) : error updating role : %v", err)
		return err
	}
	return nil
}

// Delete role
func DeleteRoleByName(name string) error {
	if err := globals.GMSDB.Delete(&Role{}, name).Error; err != nil {
		log.Printf("(Error) : error deleting role : %v", err)
		return err
	}
	return nil
}
