package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sidie88/stockbit/config"
	"net/http"

	"github.com/sidie88/stockbit/app/model"
	log "github.com/sirupsen/logrus"
)

const OmDbApiUrl string = "http://www.omdbapi.com/?apikey=%s"

type OmDbApiService interface {
	SearchWithPagination(page, searchWord string) (*model.MovieResponse, error)
	GetMovieDetail(imdbID string) (*model.MovieDetailResponse, error)
}

type OmDbApiServiceImpl struct {
	ApiKey string
}

func NewOmDbApiService(c *config.OmDbApiConfig) *OmDbApiServiceImpl {
	return &OmDbApiServiceImpl{
		ApiKey: c.ApiKey,
	}
}

func (o *OmDbApiServiceImpl) SearchWithPagination(page, searchWord string) (*model.MovieResponse, error) {
	url := fmt.Sprintf(OmDbApiUrl+"&s=%s&page=%s", o.ApiKey, searchWord, page)
	responseBody := &model.MovieResponse{}
	err := callOutbound(url, responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil

}

func (o *OmDbApiServiceImpl) GetMovieDetail(imdbID string) (*model.MovieDetailResponse, error) {
	url := fmt.Sprintf(OmDbApiUrl+"&i=%s", o.ApiKey, imdbID)
	responseBody := &model.MovieDetailResponse{}
	err := callOutbound(url, responseBody)
	if err != nil {
		return nil, err
	}
	if responseBody.ImdbID == "" {
		return nil, errors.New("response body was empty")
	}
	return responseBody, nil
}

func callOutbound(url string, responseBody model.OutboundResponse) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to call, status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(responseBody)
	if err != nil {
		return err
	}

	log.Printf("Response: %T=%s", responseBody, responseBody.ToString())

	return nil
}
