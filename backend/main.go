package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dwarvesf/devops-assignment/backend/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Close()

	if os.Getenv("DEBUG") == "1" {
		db.LogMode(true)
	}

	db.AutoMigrate(models.Accident{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	router := mux.NewRouter()
	router.HandleFunc("/accidents", func(w http.ResponseWriter, r *http.Request) {
		var a []models.Accident
		err := db.Order("date desc").Find(&a).Error
		if err != nil {
			respondWithError(w, http.StatusNotFound, "cannot find any accidents")
			return
		}
		respondWithJSON(w, http.StatusOK, a)
	}).Methods("GET")
	fmt.Println(fmt.Sprintf("Server has been started at port :%s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
