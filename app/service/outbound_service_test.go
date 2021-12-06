package service

import (
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/sidie88/stockbit/app/model"
	"reflect"
	"testing"
)

const apiKey = "12345"

func TestOmDbApiServiceImpl_GetMovieDetail(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	response := model.MovieDetailResponse{
		Title:    "Batman: The Killing Joke",
		Year:     "2016",
		Rated:    "R",
		Released: "25 Jul 2016",
		Runtime:  "76 min",
		Genre:    "Animation, Action, Crime",
		Director: "Sam Liu",
		Writer:   "Brian Azzarello, Brian Bolland, Bob Kane",
		Actors:   "Kevin Conroy, Mark Hamill, Tara Strong",
		Plot:     "As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.",
		Language: "English",
		Country:  "United States",
		Awards:   "1 win & 2 nominations",
		Poster:   "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		Ratings: []*model.Rating{
			{
				Source: "Internet Movie Database",
				Value:  "6.4/10",
			},
			{
				Source: "Rotten Tomatoes",
				Value:  "39%",
			},
		},

		Metascore:  "N/A",
		ImdbRating: "6.4",
		ImdbVotes:  "53,492",
		ImdbID:     "tt4853102",
		Type:       "movie",
		DVD:        "02 Aug 2016",
		BoxOffice:  "$3,775,000",
		Production: "N/A",
		Website:    "N/A",
		Response:   "True",
	}

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=tt12345", apiKey)
	urlUnAuthorized := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=tt12345", "apiKey")

	httpmock.RegisterResponder("GET",
		url,
		httpmock.NewJsonResponderOrPanic(200, response),
	)
	httpmock.RegisterResponder("GET",
		urlUnAuthorized,
		httpmock.NewJsonResponderOrPanic(401, nil),
	)

	type fields struct {
		ApiKey string
	}
	type args struct {
		imdbID string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "GetMovieDetail Success",
			fields:  fields{ApiKey: apiKey},
			args:    args{imdbID: "tt12345"},
			want:    &response,
			wantErr: false,
		},
		{
			name:    "GetMovieDetail Failed",
			fields:  fields{ApiKey: "apiKey"},
			args:    args{imdbID: "tt12345"},
			want:    errors.New("failed to call, status code: 401"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OmDbApiServiceImpl{
				ApiKey: tt.fields.ApiKey,
			}
			got, err := o.GetMovieDetail(tt.args.imdbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovieDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr && !reflect.DeepEqual(err, tt.want){
				t.Errorf("GetMovieDetail() error = %v, wantErr %v", err, tt.want)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMovieDetail() got = %v\n, want %v", got, tt.want)
			}
		})
	}
}

func TestOmDbApiServiceImpl_SearchWithPagination(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	response := &model.MovieResponse{
		Search:       []*model.Movie{
			{
				Title:  "Batman: The Killing Joke",
				Year:   "2016",
				ImdbID: "tt4853102",
				Type:   "movie",
				Poster: "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
			},
		},
		TotalResults: 1,
		Response:     "True",
	}

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=Batman&page=1", apiKey)

	httpmock.RegisterResponder("GET",
		url,
		httpmock.NewJsonResponderOrPanic(200, response),
	)

	urlFailed := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=Batman&page=100", apiKey)

	httpmock.RegisterResponder("GET",
		urlFailed,
		httpmock.NewJsonResponderOrPanic(500, nil),
	)

	type fields struct {
		ApiKey string
	}
	type args struct {
		page       string
		searchWord string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "SearchWithPagination Success",
			fields:  fields{ApiKey: apiKey},
			args:    args{
				page:       "1",
				searchWord: "Batman",
			},
			want:    response,
			wantErr: false,
		},
		{
			name:    "SearchWithPagination Failed",
			fields:  fields{ApiKey: apiKey},
			args:    args{
				page:       "100",
				searchWord: "Batman",
			},
			want:    errors.New("failed to call, status code: 500"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OmDbApiServiceImpl{
				ApiKey: tt.fields.ApiKey,
			}
			got, err := o.SearchWithPagination(tt.args.page, tt.args.searchWord)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchWithPagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr && !reflect.DeepEqual(err, tt.want){
				t.Errorf("SearchWithPagination() error = %v, wantErr %v", err, tt.want)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchWithPagination() got = %v, want %v", got, tt.want)
			}
		})
	}
}
