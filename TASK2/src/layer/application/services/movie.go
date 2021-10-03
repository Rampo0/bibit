package services

import (
	"github.com/rampo0/bibit/layer/domain"
	"github.com/rampo0/bibit/layer/infrastructure/rest"
	"github.com/rampo0/bibit/utils/errors"
)

type MovieService interface {
	Detail(string) (*domain.Movie, *errors.BaseErr)
	Search(string, string) (*domain.SearchMovieResponse, *errors.BaseErr)
}

type service struct {
	restRepository rest.MovieRepository
}

func NewMovieService(restRepository rest.MovieRepository) MovieService {
	return &service{
		restRepository,
	}
}

func (s *service) Detail(id string) (*domain.Movie, *errors.BaseErr) {
	return s.restRepository.Detail(id)
}

func (s *service) Search(title string, page string) (*domain.SearchMovieResponse, *errors.BaseErr) {
	return s.restRepository.Search(title, page)
}
