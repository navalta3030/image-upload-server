package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"image_upload_server/constants"
	"image_upload_server/routes"
	"image_upload_server/utils"
)

var (
	db  *gorm.DB
	err error
)

func main() {
	utils.InitializeDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/upload", routes.ImageUploadRoute).Methods("POST")
	r.HandleFunc("/getRecord", routes.GetUserImageRoute).Methods("POST")

	r.HandleFunc("/test", routes.ReadinessRoute).Methods("GET")
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Println("Listen on port 8021...")
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(constants.Port, handlers.CORS()(loggedRouter)))

}
