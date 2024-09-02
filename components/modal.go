package components

func CreateModalRegister() string {
	return `
 		<div class="modal fade" id="registerModal" tabindex="-1" aria-labelledby="registerModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="registerModalLabel">Registro de Usuario</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                            <div class="mb-3">
                                <label for="regUsername" class="form-label">Nombre de Usuario</label>
                                <input type="text" class="form-control" id="regUsername" required>
                            </div>
                            <div class="mb-3">
                                <label for="email" class="form-label">Correo Electrónico</label>
                                <input type="email" class="form-control" id="email" required>
                            </div>
                            <div class="mb-3">
                                <label for="regPassword" class="form-label">Contraseña</label>
                                <input type="password" class="form-control" id="regPassword" placeholder="Contraseña" required>
                            </div>
                            <div class="mb-3">
                                <label for="confirmPassword" class="form-label">Repetir Contraseña</label>
                                <input type="password" class="form-control" id="confirmPassword" required>
                            </div>
                            <div class="form-check mb-3">
                                <input type="checkbox" class="form-check-input" id="terms" required>
                                <label class="form-check-label" for="terms">Aceptar políticas de servicio</label>
                                <span id="errorRegAccount"></span>
                            </div>
                            <button type="submit" onclick="registerAccount()" class="btn btn-primary w-100">Registrar</button>
                    </div>
                </div>
            </div>
        </div>
	`
}

func CreateModalCreateCharacter() string {

	return `
        <div class="modal fade" id="registerCharacter" tabindex="-1" aria-labelledby="registerModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="registerModalLabel">Registro de Usuario</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                        
                            <div class="mb-3">
                                <label for="nameCharacter" class="form-label">Nombre de Personaje</label>
                                <input type="text" class="form-control" id="nameCharacter" required>
                            </div>
                            <div class="mb-3">
                                <label for="radioMale" class="form-label">Sexo del Personaje</label>
                                <div class="form-check">
                                    <input class="form-check-input" type="radio" id="radioMale">
                                    <label class="form-check-label" for="radioMale">
                                      Hombre
                                    </label>
                                  </div>
                                  <div class="form-check">
                                    <input class="form-check-input" type="radio" id="radioFemale">
                                    <label class="form-check-label" for="radioFemale">
                                      Mujer
                                    </label>
                                  </div>
                                  <span id="errorCreateCharacter"></span>
                            </div>
                            <button type="submit" class="btn btn-primary w-100" onclick="createCharacter()">Registrar Personaje</button>
                        
                    </div>
                </div>
            </div>
        </div>
	
	`
}
