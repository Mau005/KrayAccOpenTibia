package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/models"
)

func CreateServerStatus(status models.ServerStatus) (components string) {
	components = fmt.Sprintf(`
                        <ul class="custom-list">
                            <li class="custom-list-item">
                                <span class="item-message">Jugadores Online: </span>
                                <span class="item-quantity">%d</span>
                            </li>
                            <li class="custom-list-item">
                                <span class="item-message">Creaturas</span>
                                <span class="item-quantity">%d</span>
                            </li>
                            <li class="custom-list-item">
                                <span class="item-message">NPCs:</span>
                                <span class="item-quantity">%d</span>
                            </li>
                            <li class="custom-list-item">
                                <span class="item-message">Tiempo Online: </span>
                                <span id="counter" class="item-quantity">Server OFF</span>
                            </li>
                        </ul>
	`, status.Players.Online, status.Monsters.Total, status.NPCs.Total)

	return
}

func CreateRates(status models.ServerStatus) (components string) {
	components = fmt.Sprintf(`
<ul class="list-group">
                            <h4>Rates</h4>
                            <li class="list-group-item">Experiencia: %d</li>
                            <li class="list-group-item">Skills: %d</li>
                            <li class="list-group-item">Loot: %d</li>
                            <li class="list-group-item">Magic: %d</li>
                        </ul>
	`, status.Rates.Experience, status.Rates.Skill, status.Rates.Loot, status.Rates.Magic)

	return
}
