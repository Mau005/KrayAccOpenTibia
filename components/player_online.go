package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/controller"
)

func CreatePlayerOnline() string {
	var playerCtl controller.PlayerController
	itemPlayers := ""
	players := playerCtl.GetPlayerOnline()

	for _, player := range players {
		itemPlayers += fmt.Sprintf(`
		
										<tr>
                                            <td>%s</td>
                                            <td>%s</td>
                                            <td>%d</td>
                                            <td>%d</td>
                                        </tr>
		`, FunctionImagenSourcePlayer(player), player.Name, player.Level, player.Experience)
	}

	return fmt.Sprintf(`
	
		<h1>Quien esta Online?</h1>
                        <!-- Tabla de Lista de Personajes -->
                        <div class="card mb-4">
                            <div class="card-header">
                                Lista de Personajes
                            </div>
                            <div class="card-body">
                                <table class="table">
                                    <thead>
                                        <tr>
                                            <th scope="col">Outfits</th>
                                            <th scope="col">Nombre</th>
                                            <th scope="col">Nivel</th>
                                            <th scope="col">Experiencia</th>
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
