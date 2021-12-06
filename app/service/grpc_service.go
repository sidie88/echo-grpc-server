package service

import (
	"context"
	proto "github.com/sidie88/stockbit/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type GRpcService struct {
	OutboundService OmDbApiService
}

func NewGRpcService(os *OmDbApiServiceImpl) *GRpcService {
	return &GRpcService{
		OutboundService: os,
	}
}

func (g *GRpcService) SearchWithPagination(_ context.Context, request *proto.MovieRequest) (*proto.MovieResponse, error) {
	resp, err := g.OutboundService.SearchWithPagination(strconv.Itoa(int(request.Pagination)), request.Searchword)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "outbound call was failed")
	}
	var movies []*proto.Movie
	for _,v := range resp.Search {
		movie := &proto.Movie{
			Title:  v.Title,
			Year:   v.Year,
			ImdbId: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}
		movies = append(movies, movie)
	}

	totalResult, _ := strconv.Atoi(resp.TotalResults)
	movieResponse := &proto.MovieResponse{
		Search:       movies,
		TotalResults: int32(totalResult),
		Response:     resp.Response,
	}
	return movieResponse, nil
}

func (g *GRpcService) GetMovieDetail(_ context.Context, request *proto.MovieDetailId) (*proto.MovieDetailResponse, error) {
	resp, err := g.OutboundService.GetMovieDetail(request.ImdbId)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "outbound call was failed")
	}
	var ratings []*proto.Rating
	for _, v := range resp.Ratings {
		rating := &proto.Rating{
			Source: v.Source,
			Value:  v.Value,
		}
		ratings = append(ratings, rating)
	}

	movieDetailResp := &proto.MovieDetailResponse{
		Title:      resp.Title,
		Year:       resp.Year,
		Rated:      resp.Rated,
		Released:   resp.Released,
		Runtime:    resp.Runtime,
		Genre:      resp.Genre,
		Director:   resp.Director,
		Writer:     resp.Writer,
		Actors:     resp.Actors,
		Plot:       resp.Plot,
		Language:   resp.Language,
		Country:    resp.Country,
		Awards:     resp.Awards,
		Poster:     resp.Poster,
		Ratings:    ratings,
		Metascore:  resp.Metascore,
		ImdbRating: resp.ImdbRating,
		ImdbVotes:  resp.ImdbVotes,
		ImdbId:     resp.ImdbID,
		Type:       resp.Type,
		Dvd:        resp.DVD,
		BoxOffice:  resp.BoxOffice,
		Production: resp.Production,
		Website:    resp.Website,
		Response:   resp.Response,
	}

	return movieDetailResp, nil
}
