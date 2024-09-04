package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
)

func CreateTopPlayerComponent(countPlayer int) string {
	var playerCtl controller.PlayerController
	players := playerCtl.GetPlayerLimits(countPlayer)

	liItems := ""

	for value, player := range players {
		urlImage := fmt.Sprintf("%s/animoutfit.php?id=%d&addons=%d&head=%d&body=%d&legs=%d&feet=%d&mount=0&direction=3",
			config.Global.ServerWeb.UrlOutfitsView, player.LookType, player.LookAddons, player.LookHead, player.LookBody, player.LookLegs, player.LookFeet)
		liItems += fmt.Sprintf(`
<li class="list-group-item">
	<img src="%s" alt="Jugador %d"> %s - Nivel %d
</li>
		`, urlImage, value, player.Name, player.Level)
	}

	componentes := `
<h4>Top %d Jugadores</h4>
<ul class="list-group">
	%s
</ul>
	
	`

	return fmt.Sprintf(componentes, countPlayer, liItems)
}
