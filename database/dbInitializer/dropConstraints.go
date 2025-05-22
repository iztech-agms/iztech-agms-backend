package dbinitializer

import (
	"graduation-system/globals"
	"log"
)

func DropConstraints() {
	if err := globals.GMSDB.Exec(`
		ALTER TABLE users 
		DROP FOREIGN KEY fk_users_role;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from users table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE students 
		DROP FOREIGN KEY fk_students_advisor_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from students table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE advisors 
		DROP FOREIGN KEY fk_advisors_department_name;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from advisors table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE department_secretaries 
		DROP FOREIGN KEY fk_department_secretaries_department_name;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from department_secretaries table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE faculty_secretaries 
		DROP FOREIGN KEY fk_faculty_secretaries_faculty_name;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from faculty_secretaries table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE departments 
		DROP FOREIGN KEY fk_departments_faculty_name;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from departments table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE graduation_statuses 
		DROP FOREIGN KEY fk_graduation_statuses_student_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from graduation_statuses table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE notifications 
		DROP FOREIGN KEY fk_notifications_user_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from notifications table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE student_affairs 
		DROP FOREIGN KEY fk_student_affairs_user_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from student_affairs table: %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE students 
		DROP FOREIGN KEY fk_students_user_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from students table (user FK): %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE advisors 
		DROP FOREIGN KEY fk_advisors_user_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from advisors table (user FK): %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE department_secretaries 
		DROP FOREIGN KEY fk_department_secretaries_user_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from department_secretaries table (user FK): %v", err)
	}

	if err := globals.GMSDB.Exec(`
		ALTER TABLE faculty_secretaries 
		DROP FOREIGN KEY fk_faculty_secretaries_user_id;
	`).Error; err != nil {
		log.Printf("(Error) : error dropping constraint from faculty_secretaries table (user FK): %v", err)
	}

}
