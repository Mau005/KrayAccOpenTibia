package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type ExceptionController struct{}

func (ec *ExceptionController) Exeption(msg string, statusCode int, w http.ResponseWriter) {
	var exp models.Exception
	w.WriteHeader(statusCode)
	exp.Msg = msg
	exp.TimeNow = time.Now()

	err := json.NewEncoder(w).Encode(&exp)
	if err != nil {
		utils.ErrorFatal(err.Error())
	}
}

func (ec *ExceptionController) MessageAproved(msg string, w http.ResponseWriter) {
	var exp models.Exception
	w.WriteHeader(http.StatusOK)
	exp.Msg = msg
	exp.TimeNow = time.Now()
	err := json.NewEncoder(w).Encode(&exp)
	if err != nil {
		utils.ErrorFatal(err.Error())
	}
}
