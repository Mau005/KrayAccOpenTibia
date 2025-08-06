package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/models"
)

func CreateNavbar(navweb models.NavWeb) string {
	//componentsAccount := ""
	buttonRegister := ""
	if navweb.Authentication {
		// componentsAccount = `							<li><a class="dropdown-item" href="#">Cambiar contraseña</a></li>`
		buttonRegister = fmt.Sprintf(`
					<li class="nav-item">
						<a class="nav-link" href="/auth/my_board">Bienvenido %s</a>
					</li>
		`, navweb.AccountName)

	} else {
		// componentsAccount = `							<li><a class="dropdown-item" href="#">Recuperar cuenta</a></li>`
		buttonRegister = `
					<button type="button" class="btn-glow" data-bs-toggle="modal" data-bs-target="#registerModal">
						Registrarse
					</button>
		`
	}

	return fmt.Sprintf(`
 	<nav class="navbar navbar-expand-lg fixed-top">
		<div class="container-fluid">
			<a class="navbar-brand" href="#">
				<img src="/www/img/icon.png" alt="Logo">
			</a>
	
			<!-- Input para buscar personajes -->
			<form action="/get_character" method="POST" class="d-flex ms-2">
				<input class="form-control me-2" name="search" type="search" placeholder="Buscar personaje" aria-label="Buscar">
				<button class="btn btn-outline-light" type="submit">Buscar</button>
			</form>
	
			<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
				aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
			
			<div class="collapse navbar-collapse justify-content-end" id="navbarNav">
				<ul class="navbar-nav">
					<li class="nav-item">
						<a class="nav-link" href="/">Inicio</a>
					</li>

					<li class="nav-item dropdown">
						<a class="nav-link dropdown-toggle" href="#" id="navbarDropdownComunidad" role="button" data-bs-toggle="dropdown" aria-expanded="false">
							Biblioteca
						</a>
						<ul class="dropdown-menu" aria-labelledby="navbarDropdownComunidad">
							<li><a class="dropdown-item" href="#">Mapa del Mundo</a></li>
							<li><a class="dropdown-item" href="#">Task Info</a></li>
							<li><a class="dropdown-item" href="#">Descargas</a></li>
						</ul>
					</li>
					
					<!-- Menú desplegable para Comunidad -->
					<li class="nav-item dropdown">
						<a class="nav-link dropdown-toggle" href="#" id="navbarDropdownComunidad" role="button" data-bs-toggle="dropdown" aria-expanded="false">
							Comunidad
						</a>
						<ul class="dropdown-menu" aria-labelledby="navbarDropdownComunidad">
							<li><a class="dropdown-item" href="/who_online">Quien Esta Online?</a></li>
							<li><a class="dropdown-item" href="/highscore">Highscores</a></li>
							<li><a class="dropdown-item" href="/last_death">Ultimas Muertes Registradas</a></li>
						</ul>
					</li>
	


					%s
				</ul>
			</div>
		</div>
	</nav>
	
	`, buttonRegister)
}
