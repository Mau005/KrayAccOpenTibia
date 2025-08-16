package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/src/config"
	"github.com/Mau005/KrayAccOpenTibia/src/controller"
)

type ManifestHandler struct{}

func (mh *ManifestHandler) GetManifiest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&config.Manifest)
}

func (mh *ManifestHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&controller.TempData.ServStatusTotal)
}
