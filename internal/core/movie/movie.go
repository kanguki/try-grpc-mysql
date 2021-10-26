package movie

type Status int32

const (
	Status_UNKNOWN  Status = iota
	Status_ONGOING  Status = 1
	Status_FINISHED Status = 2
	Status_RUMOR    Status = 3
)

type Movie struct {
	Id        int32
	Title     string
	Director  string
	Thumbnail string
	Status    Status
	Country   string
}

func MovieToMovieRes(m Movie) *MovieRes {
	return &MovieRes{
		Id:        m.Id,
		Title:     m.Title,
		Director:  m.Director,
		Thumbnail: m.Thumbnail,
		Status:    int32(m.Status),
		Country:   m.Country,
	}
}

func MovieResToMovie(m *MovieRes) *Movie {
	return &Movie{
		Id:        m.Id,
		Title:     m.Title,
		Director:  m.Director,
		Thumbnail: m.Thumbnail,
		Status:    Status(m.Status),
		Country:   m.Country,
	}
}
