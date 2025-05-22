package dbinitializer

import (
	"graduation-system/globals"
	"log"
)

func InitConstraints() {
	DropConstraints()

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE users 
		ADD CONSTRAINT fk_users_role 
		FOREIGN KEY (role) REFERENCES roles(name) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to users table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE students 
		ADD CONSTRAINT fk_students_advisor_id 
		FOREIGN KEY (advisor_id) REFERENCES advisors(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to students table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE advisors 
		ADD CONSTRAINT fk_advisors_department_name 
		FOREIGN KEY (department_name) REFERENCES departments(department_name) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to advisors table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE department_secretaries 
		ADD CONSTRAINT fk_dep_sec_department_name 
		FOREIGN KEY (department_name) REFERENCES departments(department_name) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to department_secretaries table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE faculty_secretaries 
		ADD CONSTRAINT fk_fac_sec_faculty_name 
		FOREIGN KEY (faculty_name) REFERENCES faculties(faculty_name) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to faculty_secretaries table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE departments 
		ADD CONSTRAINT fk_departments_faculty_name 
		FOREIGN KEY (faculty_name) REFERENCES faculties(faculty_name) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to departments table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE graduation_statuses 
		ADD CONSTRAINT fk_grad_status_student_id 
		FOREIGN KEY (student_id) REFERENCES students(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to graduation_statuses table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE notifications 
		ADD CONSTRAINT fk_notifications_user_id 
		FOREIGN KEY (user_id) REFERENCES users(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to notifications table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE student_affairs 
		ADD CONSTRAINT fk_student_affairs_user_id 
		FOREIGN KEY (id) REFERENCES users(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to student_affairs table: %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE students 
		ADD CONSTRAINT fk_students_user_id 
		FOREIGN KEY (id) REFERENCES users(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to students table (user link): %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE advisors 
		ADD CONSTRAINT fk_advisors_user_id 
		FOREIGN KEY (id) REFERENCES users(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to advisors table (user link): %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE department_secretaries 
		ADD CONSTRAINT fk_dep_sec_user_id 
		FOREIGN KEY (id) REFERENCES users(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to department_secretaries table (user link): %v", err)
	}

	if err := globals.GMSDB.Exec(` 
		ALTER TABLE faculty_secretaries 
		ADD CONSTRAINT fk_fac_sec_user_id 
		FOREIGN KEY (id) REFERENCES users(id) 
		ON DELETE CASCADE;
	`).Error; err != nil {
		log.Printf("(Error) : error adding constraint to faculty_secretaries table (user link): %v", err)
	}
}
