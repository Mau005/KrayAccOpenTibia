package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"io"
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

func (ac *ApiController) PreparingCharacter(players []models.Players) []models.ClientCharacters {
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
	}
	return characters
}

func (ac *ApiController) GenerateJWT(account models.Account) (string, error) {
	expirationTime := time.Now().Add(utils.TimeSessionMinute * time.Minute)
	claims := &models.Claim{
		AccountName: account.Name,
		Email:       account.Email,
		AccountID:   account.ID,
		TypeAccess:  account.Type,
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

func (ac *ApiController) CheckOnlineServer(ip, port string) (models.ServerStatus, error) {
	var serverStatus models.ServerStatus
	packet := []byte{6, 0, 255, 255, 'i', 'n', 'f', 'o'}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), 1*time.Second)
	if err != nil {
		utils.Warn("Error connecting server:" + err.Error())
		return serverStatus, nil
	}

	defer conn.Close()

	_, err = conn.Write(packet)
	if err != nil {
		utils.Error("Error  writing packet:" + err.Error())
		return serverStatus, err
	}

	answer, err := io.ReadAll(conn)
	if err != nil {
		utils.Error("Error  reading response:" + err.Error())
		return serverStatus, err
	}

	if len(answer) == 0 {
		return serverStatus, err
	}

	err = xml.Unmarshal(answer, &serverStatus)
	if err != nil {
		utils.Error(err.Error())
		return serverStatus, err
	}

	return serverStatus, err

}
