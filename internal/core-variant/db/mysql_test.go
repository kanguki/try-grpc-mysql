package db

import (
	"testing"

	"github.com/kanguki/go-grpc-mysql/internal/core/movie"
)

var db Mysql
func TestMain(t *testing.M) {
	db = Mysql{
		ConnectString: "mo:123qwe@tcp(localhost:3306)/movie-grpc",
	}
	db.Connect()
}

func TestDbMovie(t *testing.T) {
	arr := []movie.Movie{
		{
			Id:                   1,
			Title:                "5 cm/s",
			Director:             "Nakamoto",
			Thumbnail:            "link",
			Status:               2,
			Country:              "Japan",
		},
		{
			Id:                   2,
			Title:                "Spirited away",
			Director:             "Nakamoto",
			Thumbnail:            "link",
			Status:               2,
			Country:              "Japan",
		},
		{
			Id:                   3,
			Title:                "AAA",
			Director:             "Mo",
			Thumbnail:            "link",
			Status:               3,
			Country:              "Vietnam",
		},
	}
	for i := range arr {
		err := db.AddMovie(&movie.AddMovieReq{Movie: &arr[i]} )
		if err != nil {
			t.Fatalf("error db AddMovie: %v", err)
		}
	}
	{
		res, err := db.GetAllMovies(nil)
		if err != nil {
			t.Fatalf("error db GetAllMovies: %v", err)
		}
		if len(res) < 3 {
			t.Fatalf("error db GetAllMovies: not retrieve all 3 records")
		}	
	}
	{
		res, err := db.SearchMovie(&movie.SearchMovieReq{Title: "5"})
		if err != nil {
			t.Fatalf("error db SearchMovie: %v", err)
		}
		if len(res) == 0 {
			t.Fatalf("error db SearchMovie: not act as expected")
		}
	}
	{
		err := db.UpdateMovie(&movie.UpdateMovieReq{
			Movie: &movie.Movie{
				Id: 3,
				Title: "BBB",
			},
		})
		if err != nil {
			t.Fatalf("error db UpdateMovie: %v", err)
		}
	}
	{
		err := db.DeleteMovie(&movie.DeleteMovieReq{Id: []int32{1, 2, 3}})
		if err != nil {
			t.Fatalf("error db DeleteMovie: %v", err)
		}
	}
	
}