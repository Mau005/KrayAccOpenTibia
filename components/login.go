package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/models"
)

func CreateLogin(navWeb models.NavWeb) (components string) {
	if navWeb.Authentication {
		iconStatus := `<img src="/www/img/account-status_red.gif" alt="status account"> Cuenta Gratuita`
		if navWeb.IsPremmium {
			iconStatus = `<img src="/www/img/account-status_green.gif" alt="status account">Cuenta VIP`
		}
		components += fmt.Sprintf(`
                        <ul class="list-group">
                            <li class="list-group-item">
                                %s
                            </li>
                            <li class="list-group-button" onclick="redirectMenuLogin(1)">Mi Cuenta</li>
                            <li class="list-group-button"  data-bs-toggle="modal" data-bs-target="#registerCharacter">Crear Personaje</li>
                            <li class="list-group-button" onclick="redirectMenuLogin(3)">Crear Guild</li>
                            <li class="list-group-button" onclick="redirectMenuLogin(0)">Desconectarse</li>
                        </ul>
		`, iconStatus)
	} else {
		components += `
		                    <h4>Iniciar Sesión</h4>
                        <form action="" onsubmit="loginUser(event)">
                            <div class="mb-3">
                                <label for="username" class="form-label">Usuario</label>
                                <input type="text" class="form-control" id="username" >
                            </div>
                            <div class="mb-3">
                                <label for="password" class="form-label">Contraseña</label>
                                <input type="password" class="form-control" id="password" >
                                <span id="errorLogin"></span>
                            </div>
                            
                            <button type="submit" class="btn btn-primary">Ingresar</button>     
                          </form>
		`
	}
	return
}
