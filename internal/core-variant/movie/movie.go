package movie

import (
	"context"

	"github.com/kanguki/go-grpc-mysql/internal/core/db"
	m "github.com/kanguki/go-grpc-mysql/internal/core/movie"
)

type Netflix struct {
	Db db.Database
}

func (n *Netflix) AddMovie(ctx context.Context, req *m.AddMovieReq) (*m.AddMovieRes, error) {
	err := n.Db.AddMovie(req)
	if err != nil {
		return nil, err
	}
	return &m.AddMovieRes{}, nil
}
func (n *Netflix) SearchMovie(ctx context.Context, req *m.SearchMovieReq) (*m.SearchMovieRes, error) {
	results, err := n.Db.SearchMovie(req)
	if err != nil {
		return nil, err
	}
	var res []*m.MovieRes
	for i := range results {
		res = append(res, m.MovieToMovieRes(results[i]))
	}
	return &m.SearchMovieRes{Res: res}, nil
}
func (n *Netflix) UpdateMovie(ctx context.Context, req *m.UpdateMovieReq) (*m.UpdateMovieRes, error) {
	err := n.Db.UpdateMovie(req)
	if err != nil {
		return nil, err
	}
	return &m.UpdateMovieRes{}, nil
}
func (n *Netflix) DeleteMovie(ctx context.Context, req *m.DeleteMovieReq) (*m.DeleteMovieRes, error) {
	err := n.Db.DeleteMovie(req)
	if err != nil {
		return nil, err
	}
	return &m.DeleteMovieRes{}, nil
}
func (n *Netflix) GetAllMovies(ctx context.Context, req *m.GetAllMoviesReq) (*m.GetAllMoviesRes, error) {
	results, err := n.Db.GetAllMovies(req)
	if err != nil {
		return nil, err
	}
	var res []*m.MovieRes
	for i := range results {
		res = append(res, m.MovieToMovieRes(results[i]))
	}
	return &m.GetAllMoviesRes{Res: res}, nil
}
