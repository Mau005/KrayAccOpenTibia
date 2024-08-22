package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
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
