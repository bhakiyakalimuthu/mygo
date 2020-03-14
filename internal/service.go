package internal

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var _ ArtistService = (*Service)(nil)

type ArtistService interface {
	collectMbidInfo()
	getArtistInfoFromMusicBrainz() (*MusicBraizResponse, error)
	getArtistInfoFromCoverArt()
	getArtistInfoFromDiscogs()
}
type Service struct {
	logger *zap.Logger
	client http.Client
	config Config
}



func NewService(logger *zap.Logger, client http.Client, config Config) *Service {
	return &Service{
		logger: logger,
		client: client,
		config: config,
	}
}

func (h Service) collectMbidInfo()  {

}

func (h Service) getArtistInfoFromMusicBrainz() (*MusicBraizResponse, error) {
	req,err := http.NewRequest(http.MethodGet,h.config.HostName,nil )
	if err!= nil{
		return nil, errors.New("unable to construct music brainz request")
	}
	response,err:= h.client.Do(req)
	if err!= nil{
		return nil, errors.New("failed to artist info from music brainz")
	}
	var rsp MusicBraizResponse
	err = json.NewDecoder(response.Body).Decode(&rsp)
	if err!= nil{
		return nil, errors.New("failed to decode body  from music brainz response")
	}
	defer response.Body.Close()
	return &rsp,nil
}

func (h Service) getArtistInfoFromCoverArt() {
	panic("implement me")
}

func (h Service) getArtistInfoFromDiscogs() {
	panic("implement me")
}

