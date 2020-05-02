package utils

import (
	"image_upload_server/constants"
	"image_upload_server/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

// InitializeDatabase - db connection
func InitializeDatabase() {

	db, err = gorm.Open("postgres", constants.PostgreSQLDb)

	if err != nil {
		ErrorLogger.Println(err.Error())
		ErrorLogger.Fatalln("Connection Failed to Open")
	} else {
		GeneralLogger.Println("Connection Established")
	}

	db.DropTable(&models.User{})
	db.DropTable(&models.ImageLink{})
	db.AutoMigrate(&models.ImageLink{})
	db.AutoMigrate(&models.User{})
}

// UseDb - database connection
func UseDb() *gorm.DB {
	return db
}
