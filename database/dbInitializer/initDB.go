package dbinitializer

import (
	"graduation-system/crud"
	"graduation-system/globals"
	"log"
)

func CreateTables() {
	if globals.GMSDB == nil {
		log.Fatal("Database is not initialized")
	}


	//Tables (There is probably a much better way to do this. -Bahadır)
	if err := globals.GMSDB.AutoMigrate(&crud.User{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.Advisor{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}
	
	if err := globals.GMSDB.AutoMigrate(&crud.DepartmentSecretary{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.Department{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.Faculty{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.FacultySecretary{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.GraduationStatus{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.StudentAffairs{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.Student{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}

	if err := globals.GMSDB.AutoMigrate(&crud.Notification{}); err != nil {
		log.Fatalf("(Error) : error creating tables : %v", err)
	}
}
