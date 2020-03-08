package internal

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

const (
	HTTPContentType     string = "Content-Type"
	HTTPApplicationJSON string = "application/json"
)

type Controller struct {
	logger *zap.Logger
	service *Service
}


func NewController(logger *zap.Logger, service *Service) *Controller {
	return &Controller{
		logger:  logger,
		service: service,
	}
}

func (c Controller) RegisterRoute() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/music", func(r chi.Router) {
		r.Get("/",c.GetAlbum)
	})
	return router
}

func (c Controller)  GetAlbum(w http.ResponseWriter, r *http.Request)  {
	rsp,err := c.service.getArtistInfoFromMusicBrainz()
	if err!= nil {
		http.Error(w,"failed to GetAlbum",http.StatusInternalServerError)
	}

	w.Header().Set(HTTPContentType,HTTPApplicationJSON)
	w.WriteHeader(http.StatusOK)
	if err:= json.NewEncoder(w).Encode(&rsp);err!=nil {
		http.Error(w,"failed to GetAlbum",http.StatusInternalServerError)
	}
}