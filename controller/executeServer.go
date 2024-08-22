package controller

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
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
func (esc *ExecuteServerController) StartServer(path string) (*exec.Cmd, error) {
	cmd := exec.Command(path)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	utils.Info(fmt.Sprintf("server starting on PID: %d", cmd.Process.Pid))
	esc.CMD = cmd
	return cmd, nil
}

// Monitorear el servidor
func (esc *ExecuteServerController) MonitorServer() {
	// Canal para recibir señales
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		for sig := range sigChan {
			switch sig {
			case syscall.SIGINT:
				utils.Info("Recibido SIGINT. Cerrando el servidor.")
				esc.StopServer()
			case syscall.SIGTERM:
				utils.Info("Recibido SIGTERM. Cerrando el servidor.")
				esc.StopServer()
			case syscall.SIGHUP:
				utils.Info("Recibido SIGHUP. Recargando configuración.")
				esc.ReloadConfig()
			case syscall.SIGUSR1:
				utils.Info("Recibido SIGUSR1. Guardando estado del juego.")
				esc.SaveGameState()
			}
		}
	}()

	// Esperar a que el proceso del servidor termine
	err := esc.CMD.Wait()
	if err != nil {
		log.Println("Servidor C++ terminó con error:", err)
	} else {
		log.Println("Servidor C++ terminó correctamente.")
	}
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
