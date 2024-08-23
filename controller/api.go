package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/dgrijalva/jwt-go"
)

type ApiController struct{}

func (ac *ApiController) ConvertSha1(password string) string {
	preparaing := sha1.New()
	preparaing.Write([]byte(password))
	return hex.EncodeToString(preparaing.Sum(nil))
}

func (ac *ApiController) PreparingCharacter(players []models.Player) []models.ClientCharacters {
	var characters []models.ClientCharacters

	for _, player := range players {

		// caseVocation := ""
		// switch player.Vocation {
		// case 0:
		// 	caseVocation = "None"
		// case 1:
		// 	caseVocation = "Knight"
		// default:
		// 	caseVocation = "Que te importa"
		// }
		sex := false
		if player.Sex == 1 {
			sex = true
		}
		characters = append(characters, models.ClientCharacters{
			WorldID:                          0,
			Name:                             player.Name,
			IsMale:                           sex,
			Tutorial:                         false,
			Level:                            player.Level,
			Vocation:                         player.Vocation,
			OutfitID:                         player.LookType,
			HeadColor:                        player.LookHead,
			TorsoColor:                       player.LookBody,
			LegsColor:                        player.LookLegs,
			DetailColor:                      player.LookFeet,
			AddonsFlags:                      1,
			IsHidden:                         false,
			IsTournamentParticipant:          false,
			RemainIngDailyTournamentPlayTime: 2,
		})

		// characters = append(characters, map[string]interface{}{
		// 	"worldid":                          0,
		// 	"name":                             player.Name,
		// 	"ismale":                           1,
		// 	"tutorial":                         false,
		// 	"level":                            player.Level,
		// 	"vocation":                         player.Vocation,
		// 	"outfitid":                         player.LookType,
		// 	"headcolor":                        player.LookHead,
		// 	"torsocolor":                       player.LookBody,
		// 	"legscolor":                        player.LookLegs,
		// 	"detailcolor":                      player.LookFeet,
		// 	"addonsflags":                      player.LookAddons,
		// 	"ishidden":                         0,
		// 	"istournamentparticipant":          false,
		// 	"remainingdailytournamentplaytime": 0,
		// })
		log.Println(characters)
	}
	return characters
}

func (ac *ApiController) GenerateJWT(account models.Account) (string, error) {
	expirationTime := time.Now().Add(utils.TimeSessionMinute * time.Minute)
	claims := &models.Claim{
		AccountName: account.Name,
		Email:       account.Email,
		IDAccount:   account.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(utils.PasswordSecurityDefaul))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (ac *ApiController) CheckOnlineServer(forceReload bool, ip string, port string, waitAnswerTime time.Duration) *models.ServerStatus {
	packet := []byte{6, 0, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	status := &models.ServerStatus{}

	if !forceReload {
		return status
	}

	status.IsOnline = false
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), waitAnswerTime)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return status
	}
	defer conn.Close()

	_, err = conn.Write(packet)
	if err != nil {
		fmt.Println("Error writing packet:", err)
		return status
	}

	answer, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return status
	}

	if len(answer) == 0 {
		return status
	}

	var response struct {
		XMLName xml.Name `xml:"response"`
		Players struct {
			Online int `xml:"online,attr"`
			Max    int `xml:"max,attr"`
			Peak   int `xml:"peak,attr"`
		} `xml:"players"`
		Map struct {
			Name   string `xml:"name,attr"`
			Author string `xml:"author,attr"`
			Width  int    `xml:"width,attr"`
			Height int    `xml:"height,attr"`
		} `xml:"map"`
		NPCs struct {
			Total int `xml:"total,attr"`
		} `xml:"npcs"`
		Monsters struct {
			Total int `xml:"total,attr"`
		} `xml:"monsters"`
		ServerInfo struct {
			Uptime     string `xml:"uptime,attr"`
			Location   string `xml:"location,attr"`
			URL        string `xml:"url,attr"`
			Client     string `xml:"client,attr"`
			Server     string `xml:"server,attr"`
			ServerName string `xml:"serverName,attr"`
			IP         string `xml:"ip,attr"`
		} `xml:"serverinfo"`
		MOTD string `xml:"motd"`
	}

	err = xml.Unmarshal(answer, &response)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return status
	}

	status.IsOnline = true
	status.PlayersCount = response.Players.Online
	status.PlayersMaxCount = response.Players.Max
	status.PlayersPeakCount = response.Players.Peak
	status.MapName = response.Map.Name
	status.MapAuthor = response.Map.Author
	status.MapWidth = response.Map.Width
	status.MapHeight = response.Map.Height
	status.NPCs = response.NPCs.Total
	status.Monsters = response.Monsters.Total
	status.Uptime = response.ServerInfo.Uptime
	status.Location = response.ServerInfo.Location
	status.URL = response.ServerInfo.URL
	status.Client = response.ServerInfo.Client
	status.Server = response.ServerInfo.Server
	status.ServerName = response.ServerInfo.ServerName
	status.ServerIP = response.ServerInfo.IP
	status.MOTD = response.MOTD

	return status

}
