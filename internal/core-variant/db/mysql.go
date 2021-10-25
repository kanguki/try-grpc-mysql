package db

import (
	"log"

	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"github.com/kanguki/go-grpc-mysql/internal/core/movie"
)

type Mysql struct {
	Db *gorm.DB
	ConnectString string
}

func (r *Mysql) Connect() {
	db, err := gorm.Open(mysql.Open(r.ConnectString))
	if err != nil {
		log.Fatalf("Error connect to mysql. Metadata: %v", r.ConnectString)
	}
	r.Db = db
}

func (r *Mysql) AddMovie(req *movie.AddMovieReq) (error) {
	return r.Db.Save(req.Movie).Error
}
func (r *Mysql) SearchMovie(req *movie.SearchMovieReq) (res []movie.Movie, err error) {
	qr := r.Db.Model(&movie.Movie{})
	if req.Status != 0 {
		qr.Where(&movie.Movie{Status: req.Status})
	}
	if req.Country != "" {
		qr.Where("country regexp ?", req.Country)
	}
	if req.Title != "" {
		qr.Where("title regexp ?", req.Title)
	}
	if req.Director != "" {
		qr.Where("director regexp ?", req.Director)
	}
	if req.Id != 0 {
		qr.Where(&movie.Movie{Id: req.Id})
	}
	err = qr.Find(&res).Error
	return res, err
}
func (r *Mysql) UpdateMovie(req *movie.UpdateMovieReq) (error) {
	var movieDb movie.Movie
	err := r.Db.Model(&movie.Movie{}).Where(&movie.Movie{Id: req.Movie.Id}).First(&movieDb).Error
	if err != nil {
		return err
	}
	if req.Movie.Status != 0 {
		movieDb.Status = req.Movie.Status
	}
	if req.Movie.Country != "" {
		movieDb.Country = req.Movie.Country
	}
	if req.Movie.Title != "" {
		movieDb.Title = req.Movie.Title
	}
	if req.Movie.Director != "" {
		movieDb.Director = req.Movie.Director
	}
	if req.Movie.Thumbnail != "" {
		movieDb.Thumbnail = req.Movie.Thumbnail
	}
	return r.Db.Save(&movieDb).Error
}
func (r *Mysql) DeleteMovie(req *movie.DeleteMovieReq) (error) {
	return r.Db.Model(&movie.Movie{}).Delete(&movie.Movie{}, req.Id).Error
}
func (r *Mysql) GetAllMovies(*movie.GetAllMoviesReq) (res []movie.Movie, err error) {
	err = r.Db.Model(&movie.Movie{}).Find(&res).Error
	return res, err
}