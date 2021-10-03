package rest

import (
	"encoding/json"
	"net/http"

	"github.com/rampo0/bibit/layer/domain"
	"github.com/rampo0/bibit/utils/errors"
)

type MovieRepository interface {
	Search(string, string) (*domain.SearchMovieResponse, *errors.BaseErr)
	Detail(string) (*domain.Movie, *errors.BaseErr)
}

type repository struct{}

func NewMovieRepository() MovieRepository {
	return &repository{}
}

var (
	BASE_URL   = "http://www.omdbapi.com"
	restClient = &http.Client{}
)

func (r *repository) Detail(id string) (*domain.Movie, *errors.BaseErr) {
	var err error
	request, err := http.NewRequest("GET", BASE_URL+"/?apikey=faf7e5bb&i="+id, nil)
	if err != nil {
		return nil, errors.NewInternalServerError()
	}

	response, err := restClient.Do(request)
	if err != nil {
		return nil, errors.NewInternalServerError()
	}

	defer response.Body.Close()

	var data domain.Movie
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, errors.NewInternalServerError()
	}

	if data.ImdbID == "" {
		return nil, errors.NewNotFoundError("Data not found with given id")
	}

	return &data, nil
}

func (r *repository) Search(title string, page string) (*domain.SearchMovieResponse, *errors.BaseErr) {

	var err error
	request, err := http.NewRequest("GET", BASE_URL+"/?apikey=faf7e5bb&s="+title+"&page="+page, nil)
	if err != nil {
		return nil, errors.NewInternalServerError()
	}

	response, err := restClient.Do(request)
	if err != nil {
		return nil, errors.NewInternalServerError()
	}

	defer response.Body.Close()

	var data domain.SearchMovieResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, errors.NewInternalServerError()
	}

	return &data, nil
}
