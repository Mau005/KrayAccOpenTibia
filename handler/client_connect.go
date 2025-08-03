package handler

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type HandlerClientConnect struct{}

// Define your handlers
func (hcc *HandlerClientConnect) CacheInfoHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"playersonline":        controller.TempData.ServStatusTotal.Players.Online,
		"twitchstreams":        0,
		"twitchviewer":         0,
		"gamingyoutubestreams": 0,
		"gamingyoutubeviewer":  0,
	}
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Println(err)
	}
}

func (hcc *HandlerClientConnect) EventScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// Placehcclder XML parsing, replace with your own logic
	eventList := []map[string]interface{}{} // Populate this with actual data

	response := map[string]interface{}{
		"eventlist":           eventList,
		"lastupdatetimestamp": time.Now().Unix(),
	}
	hcc.RespondJSON(w, response)
}

func (hcc *HandlerClientConnect) BoostedCreatureHandler(w http.ResponseWriter, r *http.Request) {
	var boostedCreature struct {
		RaceID int
	}
	response := map[string]interface{}{
		"boostedcreature": true,
		"raceid":          boostedCreature.RaceID,
	}
	hcc.RespondJSON(w, response)
}
func (hcc *HandlerClientConnect) GetClientIP(r *http.Request) string {
	// Si estás detrás de un proxy o load balancer
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// Puede haber múltiples IPs separadas por coma
		return strings.Split(ip, ",")[0]
	}

	// Si usas Nginx u otro proxy que usa este header
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// Valor por defecto desde la conexión
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // fallback
	}
	return ip
}
func (hcc *HandlerClientConnect) loginHandler(answerExpected models.AnswerExpected, w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "application/json")
	var PoolConnectionController controller.PoolConnectionController
	ip := hcc.GetClientIP(r)
	response, err := PoolConnectionController.CharacterLoginAccountPoolConnection(answerExpected, ip)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errorCode":    3,
			"errorMessage": "incorrect credentials",
		})
		return
	}

	if err = json.NewEncoder(w).Encode(&response); err != nil {
		utils.Warn("error encode response", err.Error())
		return
	}
	return
}

func (hcc *HandlerClientConnect) PreparingHanlderClient(w http.ResponseWriter, r *http.Request) {
	var answer models.AnswerExpected
	err := json.NewDecoder(r.Body).Decode(&answer)

	if err != nil {
		log.Println("error decode body", err)
	}

	switch answer.Type {
	case "login":
		hcc.loginHandler(answer, w, r)

	case "cacheinfo":
		hcc.CacheInfoHandler(w, r)

	}
}

// Helper functions
func (hcc *HandlerClientConnect) RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func (hcc *HandlerClientConnect) RespondError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"errorCode":    3,
		"errorMessage": msg,
	})
}
