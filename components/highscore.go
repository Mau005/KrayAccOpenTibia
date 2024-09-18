package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/config"
)

func CreateHighScore() string {
	selectData := ""
	for _, pool := range config.Global.PoolServer {
		selectData += fmt.Sprintf(
			`<option value="%s">%s</option>`, pool.World.Name, pool.World.Name)
	}
	return fmt.Sprintf(`
	
						<div class="container-fluid">
                        	<h1>HighScores</h1>
                            
                                <div class="mb-3">
                                    <label for="world" class="form-label">Selecciona un mundo:</label>
                                    <select class="form-select" id="world" name="world">
                                        %s
                                    </select>
                                </div>
                                <div class="mb-3">
                                    <label for="skills" class="form-label">Selecciona un skill y su valor: </label>
                                    <select class="form-select" id="skills" name="skills">
                                        <option value="8">Level</option>
                                        <option value="0">First</option>
                                        <option value="1">Club</option>
                                        <option value="2">Axe</option>
                                        <option value="3">Sword</option>
                                        <option value="4">Distance</option>
                                        <option value="5">Shielding</option>
                                        <option value="6">Fishing</option>
                                        <option value="7">Magic Level</option>
                                    </select>
                                </div>
                                <button type="submit" onClick="highScore()" class="btn btn-primary">Buscar</button>
                            <hr>
                        <div class="card mb-4">
                            <div class="card-header">
                                Lista de Personajes
                            </div>
                            <div class="card-body">
                                <table id="tableHighScore" class="table">
                                    <thead>
                                        <tr>
                                            <th scope="col">Nombre</th>
                                            <th scope="col">Skills:</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                        </div>

	`, selectData)
}
