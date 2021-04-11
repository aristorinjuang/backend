package movie0

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aristorinjuang/backend/02_omdapi/helpers"
	"github.com/aristorinjuang/backend/02_omdapi/utility"
)

func getWriter(writers string) []*writer {
	result := []*writer{}

	for _, w := range strings.Split(writers, ", ") {
		start := strings.Index(w, "(")
		end := strings.Index(w, ")")
		name := w
		description := ""

		if start > 0 {
			name = w[0 : start-1]
			description = w[start+1 : end]
		}

		result = append(result, &writer{
			Name:        name,
			Description: description,
		})
	}

	return result
}

func getRatings(ratings []*omdbapiRating) []*rating {
	result := []*rating{}

	for _, r := range ratings {
		var base float64 = 100
		values := strings.Split(strings.ReplaceAll(r.Value, "%", ""), "/")

		if len(values) >= 2 {
			base, _ = strconv.ParseFloat(values[1], 64)
		}

		score, _ := strconv.ParseFloat(values[0], 64)
		result = append(result, &rating{
			Source: r.Source,
			Value:  score / base,
		})
	}

	return result
}

func getCurrency(c string) *currency {
	re := regexp.MustCompile(`[\d,]`)
	noComma := strings.ReplaceAll(c, ",", "")
	symbol := re.ReplaceAllString(noComma, "")
	value, _ := strconv.ParseUint(strings.ReplaceAll(noComma, symbol, ""), 10, 64)

	return &currency{
		Symbol: symbol,
		Value:  value,
	}
}

// GetIndex gets a list of movies by a search query.
func GetIndex(host string, apiKey string, search string, page uint64) []byte {
	resp, err := http.Get(fmt.Sprintf("%s/?apiKey=%s&s=%s&page=%d",
		host,
		apiKey,
		search,
		page))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	result := omdbapi{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panic(err)
	}

	response := &utility.View{
		Response: false,
		Message:  "Incorrect IMDb ID.",
	}

	if result.Response == "True" {
		movies := []Movie{}

		for _, omdbapiMovie := range result.Search {
			year, _ := strconv.ParseUint(omdbapiMovie.Year, 10, 64)
			movie := Movie{
				Title:  omdbapiMovie.Title,
				Year:   year,
				ImdbID: omdbapiMovie.ImdbID,
				Type:   omdbapiMovie.Type,
				Poster: omdbapiMovie.Poster,
			}
			movies = append(movies, movie)
		}

		total, _ := strconv.ParseUint(result.TotalResults, 10, 64)
		lastPage := uint64(math.Ceil(float64(total) / 10))
		response = &utility.View{
			Response: true,
			Message:  "Success",
			Data:     movies,
			Pagination: utility.Pagination{
				Total:       total,
				PerPage:     10,
				CurrentPage: page,
				LastPage:    lastPage,
			},
		}
	}

	body, err = json.Marshal(response)
	if err != nil {
		log.Panic(err)
	}

	return body
}

// GetDetail gets a detail of a movie by IMDB id.
func GetDetail(host string, apiKey string, id string) []byte {
	resp, err := http.Get(fmt.Sprintf("%s/?apiKey=%s&i=%s",
		host,
		apiKey,
		id,
	))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	result := omdbapiMovieDetail{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panic(err)
	}

	response := &utility.View{
		Response: false,
		Message:  "Incorrect IMDb ID.",
	}

	if result.Response == "True" {
		year, _ := strconv.ParseUint(result.Year, 10, 64)
		released, _ := time.Parse("2 Jan 2006", result.Released)
		metascore, _ := strconv.ParseFloat(result.Metascore, 64)
		imdbRating, _ := strconv.ParseFloat(result.ImdbRating, 64)
		imdbVotes, _ := strconv.ParseUint(strings.ReplaceAll(result.ImdbVotes, ",", ""), 10, 64)
		dvd, _ := time.Parse("2 Jan 2006", result.DVD)
		movieDetail := &MovieDetail{
			Title:      result.Title,
			Year:       year,
			Rated:      result.Rated,
			Released:   released,
			Runtime:    *helpers.GetSeconds(result.Runtime),
			Genres:     strings.Split(result.Genre, ", "),
			Director:   result.Director,
			Writers:    getWriter(result.Writer),
			Actors:     strings.Split(result.Actors, ", "),
			Plot:       result.Plot,
			Languages:  strings.Split(result.Language, ", "),
			Countries:  strings.Split(result.Country, ", "),
			Awards:     result.Awards,
			Poster:     result.Poster,
			Ratings:    getRatings(result.Ratings),
			Metascore:  metascore / 100,
			ImdbRating: imdbRating / 10,
			ImdbVotes:  imdbVotes,
			ImdbID:     result.ImdbID,
			Type:       result.Type,
			DVD:        dvd,
			BoxOffice:  *getCurrency(result.BoxOffice),
			Production: strings.Split(result.Production, ", "),
			Website:    strings.ReplaceAll(result.Website, "N/A", ""),
		}

		response = &utility.View{
			Response: true,
			Message:  "Success",
			Data:     movieDetail,
		}
	}

	body, err = json.Marshal(response)
	if err != nil {
		log.Panic(err)
	}

	return body
}
