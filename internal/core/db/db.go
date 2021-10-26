package db

import (
	"github.com/kanguki/go-grpc-mysql/internal/core/movie"
)

type Database interface {
	AddMovie(*movie.AddMovieReq) error
	SearchMovie(*movie.SearchMovieReq) ([]movie.Movie, error)
	UpdateMovie(*movie.UpdateMovieReq) error
	DeleteMovie(*movie.DeleteMovieReq) error
	GetAllMovies(*movie.GetAllMoviesReq) ([]movie.Movie, error)
}
