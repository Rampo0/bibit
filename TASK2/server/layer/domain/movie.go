package domain

type Movie struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

type Movies []Movie

type MovieDetailRequest struct {
	ID string `form:"id"`
}

type SearchMovieRequest struct {
	Search string `form:"s"`
	Page   string `form:"page"`
}

type SearchMovieResponse struct {
	Search       Movies
	TotalResults string
	Response     string
}
