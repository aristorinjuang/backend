package movie0

import "time"

type writer struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type rating struct {
	Source string  `json:"source"`
	Value  float64 `json:"value"`
}

type currency struct {
	Symbol string `json:"symbol"`
	Value  uint64 `json:"value"`
}

type MovieDetail struct {
	Title      string        `json:"title"`
	Year       uint64        `json:"year"`
	Rated      string        `json:"rated"`
	Released   time.Time     `json:"released"`
	Runtime    time.Duration `json:"runtime"`
	Genres     []string      `json:"genres"`
	Director   string        `json:"director"`
	Writers    []*writer     `json:"writers"`
	Actors     []string      `json:"actors"`
	Plot       string        `json:"plot"`
	Languages  []string      `json:"languages"`
	Countries  []string      `json:"countries"`
	Awards     string        `json:"awards"`
	Poster     string        `json:"poster"`
	Ratings    []*rating     `json:"ratings"`
	Metascore  float64       `json:"metascore"`
	ImdbRating float64       `json:"imdb_rating"`
	ImdbVotes  uint64        `json:"imdb_votes"`
	ImdbID     string        `json:"imdb_id"`
	Type       string        `json:"type"`
	DVD        time.Time     `json:"dvd"`
	BoxOffice  currency      `json:"box_office"`
	Production []string      `json:"production"`
	Website    string        `json:"website"`
}

type Movie struct {
	Title  string `json:"title"`
	Year   uint64 `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type omdbapiRating struct {
	Source string
	Value  string
}

type omdbapiMovieDetail struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []*omdbapiRating
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

type omdbapiMovie struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

type omdbapi struct {
	Search       []omdbapiMovie
	TotalResults string `json:"totalResults"`
	Response     string
}
