package components

import (
	"fmt"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

func CreateMyAccount(account models.Account) string {
	premmy := account.PremiumEndsAt > uint(time.Now().Unix())
	premyStatus := "Free Account"
	if premmy {
		premyStatus = "Vip Account"
	}
	content := `
<div class="account-container">
    <div class="account-header">Bienvenido a tu cuenta %s!</div>
    <div class="account-status">
        <h5>Estado de cuenta</h5>
        <p><strong>%s</strong></p>
        <p>Your Premium Time expired at Jun 26 2024, 06:14:59 CEST.</p>
        <button class="btn btn-primary">Configurador de cuenta</button>
        <button class="btn btn-success">Comprar VIP</button>
        <button class="btn btn-danger">Desconectarse</button>
    </div>

    <div class="character-list">
        <h5>Characters</h5>
        <table class="character-table">
        <thead>
            <tr>
            <th>Name</th>
            <th>Status</th>
            </tr>
        </thead>
        <tbody>
            %s
            <!-- Agrega más filas según sea necesario -->
        </tbody>
        </table>
    </div>
</div>
	`
	var playerCtl controller.PlayerController

	playersTable := ""
	for _, player := range account.Players {
		playerWorld := playerCtl.GetNameWorld(player.Name)
		playersTable += fmt.Sprintf(`
		<tr>
			<td>%s - %s - Level %d - On %s</td>
			<td><a href="#">[Editar]</a> <a href="#">[Eliminar]</a></td>
		</tr>
			`, player.Name, FunctionGetVocation(player), player.Level, playerWorld.World)
	}

	return fmt.Sprintf(content, account.Name, premyStatus, playersTable)
}
