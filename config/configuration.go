package config

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/go-yaml/yaml"
)

var Server *controller.ExecuteServerController
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
	IP                   string `yaml:"IP"`
	Port                 uint16 `yaml:"Port"`
	Debug                bool   `yaml:"Debug"`
	ApiMode              bool   `yaml:"ApiMode"`
	LengthSecurity       int    `yaml:"LengthSecurity"`
	EnvironmentVariables bool   `yaml:"EnvironmentVariables"`
	UrlItemView          string `yaml:"UrlItemView"`
	UrlOutfitsView       string `yaml:"UrlOutfitsView"`
	TargetServer         string `yaml:"TargetServer"`
}

type MySQL struct {
	Host       string `yaml:"Host"`
	Port       uint16 `yaml:"Port"`
	UserName   string `yaml:"UserName"`
	DBPassword string `yaml:"DBPassword"`
	DataBase   string `yaml:"DataBase"`
}

type Certificate struct {
	ProtolTLS bool   `yaml:"ProtocolHTTPS"`
	Chain     string `yaml:"Chain"`
	PrivKey   string `yaml:"PrivKey"`
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
	if VarEnviroment.ServerWeb.EnvironmentVariables {
		utils.Info("MySQL variables are loaded from environment")
		VarEnviroment.DB.Host = os.Getenv("MYSQL_HOST")
		VarEnviroment.DB.Port = parseEnvUint(os.Getenv("MYSQL_PORT"), 3306)
		VarEnviroment.DB.UserName = os.Getenv("MYSQL_USER")
		VarEnviroment.DB.DBPassword = os.Getenv("MYSQL_PASSWORD")
		VarEnviroment.DB.DataBase = os.Getenv("MYSQL_DATABASE")
	} else {
		utils.Info("MySQL variables are loaded to config.yml")
	}
	utils.Info("Environment variables are added OK")

	return nil
}
func parseEnvUint(key string, defaultValue uint16) uint16 {
	if value, exists := os.LookupEnv(key); exists {
		if parsedValue, err := strconv.ParseUint(value, 10, 16); err == nil {
			return uint16(parsedValue)
		}
	}
	return defaultValue
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
