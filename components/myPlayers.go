package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

func CreateMyPlayers(navWeb models.NavWeb) string {
	itemPlayers := ""
	components := ""
	for _, player := range navWeb.MyPlayers {
		urlImage := fmt.Sprintf("%s/animoutfit.php?id=%d&addons=%d&head=%d&body=%d&legs=%d&feet=%d&mount=0&direction=3",
			config.Global.ServerWeb.UrlOutfitsView, player.LookType, player.LookAddons, player.LookHead, player.LookBody, player.LookLegs, player.LookFeet)
		itemPlayers += fmt.Sprintf(`
							<li class="list-group-item">
                                %s
                                %s - Nivel %d
                            </li>
		`, urlImage, player.Name, player.Level)

		components += itemPlayers
	}

	return fmt.Sprintf(`
	 					<ul class="list-group">
                            %s
                        </ul>
	`, components)
}
