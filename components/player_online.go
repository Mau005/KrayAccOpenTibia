package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/controller"
)

func CreatePlayerOnline() string {
	var PoolConnectionCtl controller.PoolConnectionController

	whoIsOnline := `
    <h1>Quien esta Online?</h1>
    `

	data := PoolConnectionCtl.WhoIsOnlinePoolConnection()
	contentOutput := ""
	for world, players := range data {
		content := `
        
                        <!-- Acordeón -->
                        <div class="accordion mb-4" id="%s">
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="headingOne">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
                                        %s
                                    </button>
                                </h2>
                                <div id="collapseOne" class="accordion-collapse collapse show" aria-labelledby="headingOne" data-bs-parent="#%s">
                                    <div class="accordion-body">
                                        %s
                                    </div>
                                </div>
                            </div>
                        </div>
        `
		playerList := ""

		for _, player := range players {
			playerList += fmt.Sprintf(`
		
            <tr>
                <td>%s</td>
                <td>%s</td>
                <td>%d</td>
                <td>%d</td>
            </tr>
            `, FunctionImagenSourcePlayer(player), player.Name, player.Level, player.Experience)
		}
		contentOutput += fmt.Sprintf(content, world, world, world, fmt.Sprintf(`
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
                    `, playerList))
	}

	return whoIsOnline

}
