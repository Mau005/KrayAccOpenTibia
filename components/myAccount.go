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
        <button class="btn btn-success">Comprar VIP</button>
        <button class="btn btn-danger">Desconectarse</button>
    </div>
 <hr>
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
 <hr>
 <!-- 
    <h5>Cambiar Contraseña</h5>
    <div class="account-status">
    
         <form action="/auth/change_password" method="POST">
         <input type="hidden" value="%d" name="myaccount">
            <div class="mb-3">
                <label for="regPassword" class="form-label">Contraseña</label>
                <input type="password" class="form-control" name="regPassword" placeholder="Contraseña" required>
            </div>
            <div class="mb-3">
                <label for="confirmPassword" class="form-label">Repetir Contraseña</label>
                <input type="password" class="form-control" name="confirmPassword" required>
            </div>   
            <button type="submit" class="btn btn-success">Cambiar</button>
        </form>
    
    </div>
    -->

</div>
	`
	var playerCtl controller.PlayerController

	playersTable := ""
	for _, player := range account.Players {
		playerWorld := playerCtl.GetNameWorld(player.Name)
		playersTable += fmt.Sprintf(`
		<tr>
			<td><a href="/get_character/%s">%s</a> %s - Level %d - On %s</td>
			<!--<td><a href="#">[Editar]</a> <a href="#">[Eliminar]</a></td>-->
		</tr>
			`, player.Name, player.Name, FunctionGetVocation(player), player.Level, playerWorld.World)
	}

	return fmt.Sprintf(content, account.Name, premyStatus, playersTable, account.ID)
}

func CreateLogin(navWeb models.NavWeb) (components string) {
	if navWeb.Authentication {
		iconStatus := `<img src="/www/img/account-status_red.gif" alt="status account"> Cuenta Gratuita`
		if navWeb.IsPremmium {
			iconStatus = `<img src="/www/img/account-status_green.gif" alt="status account">Cuenta VIP`
		}
		components += fmt.Sprintf(`
                        <h4>Cuenta</h4>
                        <ul class="list-group">
                            <li class="list-group-item">
                                %s
                            </li>
                            <li class="list-group-button" onclick="redirectMenuLogin(1)">Mi Cuenta</li>
                            <li class="list-group-button"  data-bs-toggle="modal" data-bs-target="#registerCharacter">Crear Personaje</li>
                            <li class="list-group-button" onclick="redirectMenuLogin(0)">Desconectarse</li>
                        </ul>
		`, iconStatus)
	} else {
		components += `
		                    <h4>Iniciar Sesión</h4>
                        <form action="#" onsubmit="loginUser(event)">
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
