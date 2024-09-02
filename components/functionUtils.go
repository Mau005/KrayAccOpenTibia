package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

func FunctionImagenSourcePlayer(player models.Players) string {
	urlImg := fmt.Sprintf("%s/animoutfit.php?id=%d&addons=%d&head=%d&body=%d&legs=%d&feet=%d&mount=0&direction=3",
		config.VarEnviroment.ServerWeb.UrlOutfitsView, player.LookType, player.LookAddons, player.LookHead, player.LookBody, player.LookLegs, player.LookFeet)
	return fmt.Sprintf(`
	<img src="%s" alt="Jugador %d">
	`, urlImg, player.ID)
}
