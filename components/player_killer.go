package components

import (
	"fmt"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/controller"
)

func CreateLastPlayerKills() string {
	var playerCtl controller.PlayerController
	itemPlayers := ""
	players := playerCtl.GetPlayerDeath()

	for index, death := range players {
		time := time.Unix(int64(death.Time), 0)
		itemPlayers += fmt.Sprintf(`
		
										<tr>
                                            <td>%d</td>
                                            <td>%s</td>
                                            <td>%s</td>
                                        </tr>
		`, index+1, fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second()), fmt.Sprintf("%s died at level %d by a %s.", death.Player.Name, death.Level, death.KilledBy))
	}

	return fmt.Sprintf(`
	
		<h1>Quien esta Online?</h1>
                        <!-- Tabla de Lista de Personajes -->
                        <div class="card mb-4">
                            <div class="card-header">
                                Ultimas Muertes!
                            </div>
                            <div class="card-body">
                                <table class="table">
                                    <thead>
                                        <tr>
                                            <th scope="col">N°</th>
                                            <th scope="col">Fecha</th>
                                            <th scope="col">Descripción</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        %s
                                    </tbody>
                                </table>
                            </div>
                        </div>
	`, itemPlayers)

}
