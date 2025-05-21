package dbinitializer

import (
	"graduation-system/globals"
	"log"
)

func DropConstraints() {
	if err := globals.GMSDB.Exec(`
		ALTER TABLE advisors 
		DROP FOREIGN KEY fk_advisors_department_name;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraints from advisors table: %v", err)
	}
}
