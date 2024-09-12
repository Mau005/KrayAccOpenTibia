package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

func FunctionImagenSourcePlayer(player models.Players) string {
	urlImg := fmt.Sprintf("%s/animoutfit.php?id=%d&addons=%d&head=%d&body=%d&legs=%d&feet=%d&mount=0&direction=3",
		config.Global.ServerWeb.UrlOutfitsView, player.LookType, player.LookAddons, player.LookHead, player.LookBody, player.LookLegs, player.LookFeet)
	return fmt.Sprintf(`
	<img src="%s" alt="Jugador %d">
	`, urlImg, player.ID)
}

func FunctionGetVocation(player models.Players) string {

	switch player.Vocation {
	case 0:
		return "No Vocation"
	case 1:
		return "Sorcerer"
	case 2:
		return "Druid"
	case 3:
		return "Paladin"
	case 4:
		return "Knight"
	case 5:
		return "Master Sorcerer"
	case 6:
		return "Elder Druid"
	case 7:
		return "Royal Paladin"
	case 8:
		return "Elite Knight"
	default:
		return "No encontrado"

	}
}
