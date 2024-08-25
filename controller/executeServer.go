package controller

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type ExecuteServerController struct {
	PathServer string
	CMD        *exec.Cmd
}

func NewExecuteServerController(path string) (execute ExecuteServerController) {
	execute.PathServer = path
	return
}

// Iniciar el servidor
func (esc *ExecuteServerController) StartServer() error {
	if esc.PathServer == "" {
		return errors.New("server not active")
	}
	checkOS := runtime.GOOS
	targetExecute := ""
	targetPath := ""
	switch checkOS {

	case "windows":
		slicePath := strings.Split(esc.PathServer, "\\")
		targetExecute = fmt.Sprintf("%s.exe", slicePath[len(slicePath)-1])
		targetPath = strings.Join(slicePath[:len(slicePath)-1], "\\")

	default:
		slicePath := strings.Split(esc.PathServer, "/")
		targetExecute = fmt.Sprintf("./%s", slicePath[len(slicePath)-1])
		targetPath = strings.Join(slicePath[:len(slicePath)-1], "/")
	}

	cmd := exec.Command(fmt.Sprintf("./%s", targetExecute))
	cmd.Dir = targetPath

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	utils.Info(fmt.Sprintf("server starting on PID: %d", cmd.Process.Pid))

	esc.CMD = cmd
	return nil
}

// Cerrar el servidor
func (esc *ExecuteServerController) StopServer() error {
	if esc.CMD.Process != nil {
		err := esc.CMD.Process.Signal(syscall.SIGTERM) // Enviar señal de terminación
		if err != nil {
			return err
		}
		log.Println("Servidor C++ detenido.")
	}
	return nil
}

// Recargar la configuración
func (esc *ExecuteServerController) ReloadConfig() {
	// Implementa la lógica para recargar la configuración aquí
	log.Println("Recargando configuración...")
}

// Guardar el estado del juego
func (esc *ExecuteServerController) SaveGameState() {
	// Implementa la lógica para guardar el estado del juego aquí
	log.Println("Guardando estado del juego...")
}
