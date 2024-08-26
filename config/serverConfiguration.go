package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type ExperienceStageLua struct {
	MinLevel   uint16
	MaxLevel   uint16
	Multiplier uint16
}

type RateServerLua struct {
	RateExp   uint16
	RateSkill uint16
	RateLoot  uint16
	RateMagic uint16
	RateSpawn uint16
}

type ServerLua struct {
	IPServer           string
	LoginProtocolPort  uint16
	GameProtocolPort   uint16
	StatusProtocolPort uint16
	HouseRentPeriod    string
	WorldType          uint8
}

type ExecuteServer struct {
	NameServer      string
	NameExecute     string
	PathServer      string
	Server          ServerLua
	MySQL           MySQL
	ExperienceStage ExperienceStageLua
	RateServer      RateServerLua
	CMD             *exec.Cmd
}

// Iniciar el servidor
func (esc *ExecuteServer) StartServer() error {

	cmd := exec.Command(fmt.Sprintf("./%s", esc.NameExecute))
	cmd.Dir = esc.PathServer

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
func (esc *ExecuteServer) StopServer() error {
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
func (esc *ExecuteServer) ReloadConfig() {
	// Implementa la lógica para recargar la configuración aquí
	log.Println("Recargando configuración...")
}

// Guardar el estado del juego
func (esc *ExecuteServer) SaveGameState() {
	// Implementa la lógica para guardar el estado del juego aquí
	log.Println("Guardando estado del juego...")
}
