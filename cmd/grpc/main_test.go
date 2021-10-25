package main

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/kanguki/go-grpc-mysql/internal/core-variant/db"
	"github.com/kanguki/go-grpc-mysql/internal/core-variant/movie"
	m "github.com/kanguki/go-grpc-mysql/internal/core/movie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	db := db.Mysql{
		ConnectString: "mo:123qwe@tcp(localhost:3306)/movie-grpc",
	}
	db.Connect()

	movieService := movie.Netflix{Db: &db}
	s := grpc.NewServer()
	m.RegisterMovieServiceServer(s, &movieService)

	lis = bufconn.Listen(bufSize)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGrpServer(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "abc", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial abc: %v", err)
	}
	defer conn.Close()

	c := m.NewMovieServiceClient(conn)

	movie := m.Movie{
		Id:        1,
		Title:     "Test",
		Director:  "Mo",
		Thumbnail: "link",
		Status:    1,
		Country:   "America",
	}
	ctx := context.Background()
	{
		_, err = c.AddMovie(ctx, &m.AddMovieReq{Movie: &movie})
		if err != nil {
			t.Fatalf("error when calling AddMovie: %s", err)
		}
	}
	{
		res, err := c.SearchMovie(ctx, &m.SearchMovieReq{Title: "test"})
		if err != nil {
			t.Fatalf("error when calling SearchMovie: %s", err)
		}
		if len(res.Res) < 1 {
			t.Fatalf("error when calling SearchMovie: not act as expected")
		}
	}
	{
		movie.Country = "Hybe"
		_, err = c.UpdateMovie(ctx, &m.UpdateMovieReq{Movie: &movie})
		if err != nil {
			t.Fatalf("error when calling AddMovie: %s", err)
		}
	}
	{
		res, err := c.GetAllMovies(ctx, &m.GetAllMoviesReq{})
		if err != nil {
			t.Fatalf("error when calling GetAllMovies: %s", err)
		}
		if len(res.Res) < 1 {
			t.Fatalf("error when calling GetAllMovies: not act as expected")
		}
	}
	{
		_, err := c.DeleteMovie(ctx, &m.DeleteMovieReq{Id: []int32{1}})
		if err != nil {
			t.Fatalf("error when calling DeleteMovie: %s", err)
		}
	}
}
