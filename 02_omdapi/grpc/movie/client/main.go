package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/aristorinjuang/backend/02_omdapi/grpc/movie"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := grpc.Dial(os.Getenv("SERVER"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	c := movie.NewMovieServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	search := flag.String("search", "", "the search query")
	page := flag.Uint64("page", 1, "the page number")
	id := flag.String("id", "", "the imdb id")

	flag.Parse()

	if *search != "" {
		r, err := c.GetMovies(ctx, &movie.MoviesRequest{
			Search: *search,
			Page:   *page,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(r)
	}

	if *id != "" {
		r, err := c.GetMovie(ctx, &movie.MovieRequest{
			Id: *id,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(r)
	}
}
