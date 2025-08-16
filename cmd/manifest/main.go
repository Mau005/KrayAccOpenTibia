package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Mau005/KrayAccOpenTibia/src/config"
	"github.com/Mau005/KrayAccOpenTibia/src/controller"
	"github.com/Mau005/KrayAccOpenTibia/src/models"
)

func main() {
	// Carpeta local donde tienes los archivos del cliente
	err := config.Load("config.yml")
	if err != nil {
		log.Println(err)
		return
	}
	root := "./Client"

	// Manifest inicial
	var m models.Manifest
	m.App = "AinhoOT"
	m.Version = os.Getenv("VERSION")
	m.BaseURL = fmt.Sprintf("http://%s:%d/launcher_client", config.Global.ServerWeb.IP, config.Global.ServerWeb.Port) // Nueva ruta base

	var lauch controller.LaucherController

	// Recorremos todos los archivos de la carpeta root
	err = filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
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

		// Obtener ruta relativa para guardar en el manifest
		rel, _ := filepath.Rel(root, p)
		m.Files = append(m.Files, models.FileEntry{
			Path:   filepath.ToSlash(rel),
			Size:   size,
			SHA256: sum,
		})

		return nil
	})
	if err != nil {
		fmt.Println("Error al recorrer archivos:", err)
		os.Exit(1)
	}

	// Exportar manifest.json
	out, _ := json.MarshalIndent(m, "", "  ")
	if err := os.WriteFile("manifest.json", out, 0644); err != nil {
		fmt.Println("Error al guardar manifest.json:", err)
		os.Exit(1)
	}

	fmt.Println("Manifest generado correctamente en manifest.json")
}
