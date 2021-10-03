package services_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/rampo0/bibit/layer/application/services"
	"github.com/rampo0/bibit/layer/infrastructure/db"
	"github.com/rampo0/bibit/layer/infrastructure/rest"
)

var (
	service = services.NewMovieService(rest.NewMovieRepository(), db.NewLoggerRepository())
)

func TestSuccessDetail(t *testing.T) {
	movieId := "tt4853102"
	_, err := service.Detail(movieId)
	if err != nil {
		t.Error(err)
	}
}

func TestNotFoundDetail(t *testing.T) {
	movieId := "randomId"
	_, err := service.Detail(movieId)
	if err == nil {
		t.Errorf("Should contain rest errors")
	}

	if err.Status != http.StatusNotFound {
		t.Errorf("Should get 404 response code")
	}
}

func TestSuccessSearch(t *testing.T) {
	keyword := "Batman"
	page := "2"
	res, err := service.Search(keyword, page)
	if err != nil {
		t.Error(err)
	}
	haveResult, _ := strconv.Atoi(res.TotalResults)
	if haveResult <= 0 {
		t.Errorf("Should have result")
	}

}

func TestFailedSearch(t *testing.T) {
	keyword := "Batman"
	page := "99999999"
	res, err := service.Search(keyword, page)
	if err != nil {
		t.Error(err)
	}

	if res.Response != "False" {
		t.Errorf("Should have Response : \"False\"")
	}
}
