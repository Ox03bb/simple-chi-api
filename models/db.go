package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// type DataBase struct {
// 	DB *gorm.DB
// }

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		User{},
	)
	return err
}

func createRootUser(db *gorm.DB) error {
	var isAdmin bool = true

	var root User = User{
		Username: "root",
		Password: "root",
		Email:    "root@root.su",
		IsAdmin:  &isAdmin,
	}
	err := db.Create(&root).Error
	if err != nil {
		log.Println("Error creating root user:", err)
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func InitDB() (*gorm.DB, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}

	log.Println("Connected to the database successfully.")

	err = Migrate(db)
	if err != nil {
		log.Fatal("migration failed:", err)
		return nil, err
	}

	log.Println("Migration successful.")

	createRootUser(db)
	if err != nil {
		log.Println("Error creating root user:", err)
		return nil, err
	}
	log.Println("Root user created successfully.")

	return db, nil
}
