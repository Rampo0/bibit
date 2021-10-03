package services

import (
	"fmt"

	"github.com/rampo0/bibit/layer/domain"
	"github.com/rampo0/bibit/layer/infrastructure/db"
	"github.com/rampo0/bibit/layer/infrastructure/rest"
	"github.com/rampo0/bibit/utils/errors"
)

type MovieService interface {
	Detail(string) (*domain.Movie, *errors.BaseErr)
	Search(string, string) (*domain.SearchMovieResponse, *errors.BaseErr)
}

type movieService struct {
	restRepository rest.MovieRepository
	dbRepository   db.LoggerRepository
}

func NewMovieService(restRepository rest.MovieRepository, dbRepository db.LoggerRepository) MovieService {
	return &movieService{
		restRepository,
		dbRepository,
	}
}

func (s *movieService) Detail(id string) (*domain.Movie, *errors.BaseErr) {
	return s.restRepository.Detail(id)
}

func (s *movieService) Search(title string, page string) (*domain.SearchMovieResponse, *errors.BaseErr) {
	s.dbRepository.Log(fmt.Sprintf("Call Search API with search: %s, page: %s", title, page))
	return s.restRepository.Search(title, page)
}
