package rpc

import (
	"context"

	"github.com/rampo0/bibit/layer/application/services"
	"github.com/rampo0/bibit/layer/presentation/rpc/moviepb"
	"github.com/rampo0/bibit/utils/errors"
)

type MovieHandler interface {
	moviepb.MovieServiceServer
}

type movieHandler struct {
	service services.MovieService
	moviepb.UnimplementedMovieServiceServer
}

func NewMovieHandler(service services.MovieService) MovieHandler {
	return &movieHandler{
		service,
		moviepb.UnimplementedMovieServiceServer{},
	}
}

func (h *movieHandler) Detail(ctx context.Context, req *moviepb.DetailRequest) (*moviepb.DetailResponse, error) {
	movieId := req.GetId()
	res, err := h.service.Detail(movieId)

	if err != nil {
		return nil, errors.NewRPCErr(err.Error)
	}

	movie := &moviepb.Movie{
		Title:  res.Title,
		Year:   res.Year,
		ImdbID: res.ImdbID,
		Poster: res.Poster,
		Type:   res.Type,
	}
	response := &moviepb.DetailResponse{
		Movie: movie,
	}

	return response, nil
}

func (h *movieHandler) Search(ctx context.Context, req *moviepb.SearchRequest) (*moviepb.SearchResponse, error) {
	title := req.GetSearch()
	page := req.GetPage()

	res, err := h.service.Search(title, page)

	if err != nil {
		return nil, errors.NewRPCErr(err.Error)
	}

	var movies []*moviepb.Movie
	for _, movie := range res.Search {
		movies = append(movies, &moviepb.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			ImdbID: movie.ImdbID,
			Type:   movie.Type,
			Poster: movie.Poster,
		})
	}

	response := &moviepb.SearchResponse{
		Search:       movies,
		TotalResults: res.TotalResults,
		Response:     res.Response,
	}

	return response, nil
}
