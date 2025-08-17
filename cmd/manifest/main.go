package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Mau005/KrayAccOpenTibia/src/config"
	"github.com/Mau005/KrayAccOpenTibia/src/controller"
	"github.com/Mau005/KrayAccOpenTibia/src/handler"
	"github.com/Mau005/KrayAccOpenTibia/src/middleware"
	"github.com/Mau005/KrayAccOpenTibia/src/models"
	"github.com/Mau005/KrayAccOpenTibia/src/utils"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

func create_manifest() {

	root := "./Client"

	var m models.Manifest
	m.App = "AinhoOT"
	m.Version = os.Getenv("VERSION")
	if m.Version == "" {
		m.Version = "1.0.0"
	}
	m.BaseURL = fmt.Sprintf("http://%s:%d/launcher_client", config.Global.ServerWeb.IP, config.Global.ServerWeb.Port)
	var lauch controller.LaucherController

	err := filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		size, sum, err := lauch.HashFile(p)
		if err != nil {
			return fmt.Errorf("error al hashear %s: %w", p, err)
		}

		rel, _ := filepath.Rel(root, p)
		m.Files = append(m.Files, models.FileEntry{
			Path:   filepath.ToSlash(rel),
			Size:   size,
			SHA256: sum,
		})

		return nil
	})
	if err != nil {
		utils.ErrorFatal("Error while browsing files", err.Error())
		return
	}

	// Exportar manifest.json
	out, _ := json.MarshalIndent(m, "", "  ")
	if err := os.WriteFile("manifest.json", out, 0644); err != nil {
		utils.ErrorFatal("error save manifest.json:", err.Error())
		return
	}

	utils.InfoBlue("Manifest generated successfully in manifest.json")
}

func loadDependency() error {
	global := &models.Configuration{}
	content, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &global)
	if err != nil {
		return err
	}
	config.Global = global

	man, err := os.ReadFile("manifest.json")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(man, &config.Manifest)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := loadDependency()
	if err != nil {
		utils.ErrorFatal(err.Error())
	}

	preparing := flag.Bool("create_manifest", false, "Create manifest to client ")
	flag.Parse()
	if *preparing {
		create_manifest()
		return
	}
	controller.InitTemporaryEmpty()

	r := mux.NewRouter()
	lauch := http.FileServer(http.Dir("./client"))
	r.PathPrefix("/launcher_client").Handler(http.StripPrefix("/launcher_client", lauch))

	r.HandleFunc("/get_news_short", func(w http.ResponseWriter, r *http.Request) {
		var listNews []models.NewsShort
		listNews = append(listNews, models.NewsShort{
			IconID:      1,
			Description: "Test News Example",
		})

		json.NewEncoder(w).Encode(struct {
			NewsShort []models.NewsShort
		}{
			NewsShort: listNews,
		})
	}).Methods("GET")

	ctl := r.PathPrefix("/client").Subrouter()
	ctl.Use(middleware.CommonMiddleware)

	var maniHandler handler.ManifestHandler
	ctl.HandleFunc("/manifest", maniHandler.GetManifiest).Methods("GET")
	ctl.HandleFunc("/info", maniHandler.GetInfo).Methods("GET")

	ip := fmt.Sprintf("%s:%d", config.Global.ServerWeb.IP, config.Global.ServerWeb.Port)

	utils.InfoBlue(fmt.Sprintf("[HTTP] Starting the HTTP server: http://%s/", ip))
	if err := http.ListenAndServe(ip, r); err != nil {
		utils.ErrorFatal("Error starting HTTP server: " + err.Error())
	}

}
