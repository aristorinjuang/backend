package movie0

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aristorinjuang/backend/02_omdapi/utility"
	"github.com/gorilla/mux"
)

// SetLog logs the request and response to the database.
func SetLog(db *sql.DB, search string, page uint64, response *[]byte) {
	log.Println(search, page, string(*response))

	if _, err := db.Exec(
		"INSERT INTO logs(searchword, pagination, response, created_at) VALUES(?, ?, ?, ?)",
		search,
		page,
		response,
		time.Now(),
	); err != nil {
		log.Println(err)
	}
}

func getNoResponse() []byte {
	response := &utility.View{
		Response: false,
		Message:  "Incorrect IMDb ID.",
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Panic(err)
	}

	return body
}

// Index is a list of movies in routing.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	search := r.URL.Query().Get("searchword")
	page, _ := strconv.ParseUint(r.URL.Query().Get("pagination"), 10, 64)

	if page == 0 {
		page = 1
	}

	if search == "" {
		w.Write(getNoResponse())
		return
	}

	response := GetIndex(os.Getenv("OMDBAPI"), os.Getenv("API_KEY"), search, page)
	go SetLog(utility.DB, search, page, &response)

	w.Write(response)
}

// Show shows a detail of the movie by IMDB id in routing.
func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	if vars["id"] == "" {
		w.Write(getNoResponse())
		return
	}

	w.Write(GetDetail(os.Getenv("OMDBAPI"), os.Getenv("API_KEY"), vars["id"]))
}
