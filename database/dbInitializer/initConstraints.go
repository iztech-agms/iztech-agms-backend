package dbinitializer

import (
	"graduation-system/globals"
	"log"
)

func InitConstraints() {
	DropConstraints()

	if err := globals.GMSDB.Exec(`
		ALTER TABLE advisors 
		ADD CONSTRAINT fk_advisors_department_name FOREIGN KEY (department_name) REFERENCES departments(name) ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraints to advisors table: %v", err)
	}
}
