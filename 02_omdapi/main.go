package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	movie0 "github.com/aristorinjuang/backend/02_omdapi/movie/v0"
	"github.com/aristorinjuang/backend/02_omdapi/utility"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var err error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	utility.DB, err = sql.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v0/", movie0.Index)
	r.HandleFunc("/v0/{id}", movie0.Show)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
