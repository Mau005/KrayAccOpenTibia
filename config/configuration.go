package config

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/go-yaml/yaml"
)

var VarEnviroment *Configuration
var SecretPassword []byte
var Welcome string = `
____  __.                      _____                
|    |/ _|___________  ___.__. /  _  \   ____  ____  
|      < \_  __ \__  \<   |  |/  /_\  \_/ ___\/ ___\ 
|    |  \ |  | \// __ \\___  /    |    \  \__\  \___ 
|____|__ \|__|  (____  / ____\____|__  /\___  >___  >
	   \/           \/\/            \/     \/    \/ 
   `

type ServerWeb struct {
	IP             string `yaml:"IP"`
	Port           string `yaml:"Port"`
	Debug          bool   `yaml:"Debug"`
	ApiMode        bool   `yaml:"ApiMode"`
	LengthSecurity int    `yaml:"LengthSecurity"`
}

type MySQL struct {
	Host       string `yaml:"Host"`
	Port       uint   `yaml:"Port"`
	UserName   string `yaml:"UserName"`
	DBPassword string `yaml:"DBPassword"`
	DataBase   string `yaml:"DataBase"`
}

type Certificate struct {
	ProtocolHTTPS bool   `yaml:"ProtocolHTTPS"`
	Chain         string `yaml:"Chain"`
	PrivKey       string `yaml:"PrivKey"`
}

type Configuration struct {
	DB          MySQL       `yaml:"MySQL"`
	ServerWeb   ServerWeb   `yaml:"ServerWeb"`
	Certificate Certificate `yaml:"Certificate"`
}

func Load(filename string) error {

	VarEnviroment = &Configuration{}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &VarEnviroment)
	if err != nil {
		return err
	}

	utils.InfoBlue(Welcome)

	SecretPassword = []byte(GenerateRandomPassword(VarEnviroment.ServerWeb.LengthSecurity))

	//Cargo la configuracion
	if VarEnviroment.ServerWeb.Debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetPrefix("DEBUG: ")
		SecretPassword = []byte("debugging")
		utils.Warn("Security mode Debug")
	} else {
		utils.Info(fmt.Sprintf("Security Active %d\n", VarEnviroment.ServerWeb.LengthSecurity))
	}
	utils.Info("Environment variables are added OK")

	return nil
}

func generateRandomCharacter(charset string) string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		log.Fatal(err)
	}
	return string(charset[n.Int64()])
}

func GenerateRandomPassword(length int) string {
	allChars := utils.UpperCase + utils.LowerCase + utils.Digits + utils.Special
	password := ""
	for i := 0; i < length; i++ {
		randomChar := generateRandomCharacter(allChars)
		password += randomChar
	}
	return password
}
