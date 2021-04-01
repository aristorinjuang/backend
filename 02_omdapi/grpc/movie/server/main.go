package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/aristorinjuang/backend/02_omdapi/grpc/movie"
	movie0 "github.com/aristorinjuang/backend/02_omdapi/movie/v0"
	"github.com/aristorinjuang/backend/02_omdapi/utility"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	movie.UnimplementedMovieServiceServer
}

var (
	db  *sql.DB
	err error
)

// GetMovies returns a list of movies as a proto response.
func (s *server) GetMovies(ctx context.Context, in *movie.MoviesRequest) (*movie.MoviesResponse, error) {
	body := movie0.GetIndex(os.Getenv("OMDBAPI"), os.Getenv("API_KEY"), in.GetSearch(), in.GetPage())
	response := &utility.View{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Panic(err)
	}

	result := &movie.MoviesResponse{
		Response: response.Response,
		Message:  response.Message,
		Pagination: &movie.Pagination{
			Total:       response.Pagination.Total,
			PerPage:     response.Pagination.PerPage,
			CurrentPage: response.Pagination.CurrentPage,
			LastPage:    response.Pagination.LastPage,
		},
	}

	moviesBytes, _ := json.Marshal(response.Data)
	if err := json.Unmarshal(moviesBytes, &result.Data); err != nil {
		log.Panic(err)
	}

	resultBytes, _ := json.Marshal(result)
	go movie0.SetLog(db, in.GetSearch(), in.GetPage(), &resultBytes)

	return result, nil
}

// GetMovie returns a detail of a movie as a proto response.
func (s *server) GetMovie(ctx context.Context, in *movie.MovieRequest) (*movie.MovieResponse, error) {
	body := movie0.GetDetail(os.Getenv("OMDBAPI"), os.Getenv("API_KEY"), in.GetId())
	response := &utility.View{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Panic(err)
	}

	movieDetailBytes, _ := json.Marshal(response.Data)
	movieDetail := &movie0.MovieDetail{}
	if err := json.Unmarshal(movieDetailBytes, &movieDetail); err != nil {
		log.Panic(err)
	}

	writers := []*movie.Writer{}
	for _, writer := range movieDetail.Writers {
		writers = append(writers, &movie.Writer{
			Name:        writer.Name,
			Description: writer.Description,
		})
	}

	ratings := []*movie.Rating{}
	for _, rating := range movieDetail.Ratings {
		ratings = append(ratings, &movie.Rating{
			Source: rating.Source,
			Value:  float32(rating.Value),
		})
	}

	result := &movie.MovieResponse{
		Response: response.Response,
		Message:  response.Message,
		Data: &movie.MovieDetail{
			Title:      movieDetail.Title,
			Year:       movieDetail.Year,
			Rated:      movieDetail.Rated,
			Released:   timestamppb.New(movieDetail.Released),
			Runtime:    durationpb.New(movieDetail.Runtime),
			Genres:     movieDetail.Genres,
			Director:   movieDetail.Director,
			Writers:    writers,
			Actors:     movieDetail.Actors,
			Plot:       movieDetail.Plot,
			Languages:  movieDetail.Languages,
			Countries:  movieDetail.Countries,
			Awards:     movieDetail.Awards,
			Poster:     movieDetail.Poster,
			Ratings:    ratings,
			Metascore:  float32(movieDetail.Metascore),
			ImdbRating: float32(movieDetail.ImdbRating),
			ImdbVotes:  float32(movieDetail.ImdbVotes),
			ImdbId:     movieDetail.ImdbID,
			Type:       movieDetail.Type,
			Dvd:        timestamppb.New(movieDetail.DVD),
			BoxOffice: &movie.Currency{
				Symbol: movieDetail.BoxOffice.Symbol,
				Value:  movieDetail.BoxOffice.Value,
			},
			Production: movieDetail.Production,
			Website:    movieDetail.Website,
		},
		Pagination: &movie.Pagination{
			Total:       response.Pagination.Total,
			PerPage:     response.Pagination.PerPage,
			CurrentPage: response.Pagination.CurrentPage,
			LastPage:    response.Pagination.LastPage,
		},
	}

	return result, nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Panic(err)
	}
	s := grpc.NewServer()
	movie.RegisterMovieServiceServer(s, &server{})
	if err := s.Serve(ln); err != nil {
		log.Fatal(err)
	}
}
