function createCharacter() {
    let nameCharacter = document.getElementById("nameCharacter").value;
    let isMale = document.getElementById("radioMale").checked;
    let isFemale = document.getElementById("radioFemale").checked;

    const errorCreateCharacter = document.getElementById("errorCreateCharacter");

    if (isMale == false && isFemale == false) {
        const errorElement = document.getElementById("errorCreateCharacter");
        errorElement.innerHTML = "Tiene que indicar que sexo tiene";
        errorElement.style.color = "red"; // Cambia el color a rojo
        return
    }
    let sex = 0; 
    if (isMale) {
        sex = 1;
    }

    sendRequest("/auth/create_character", "POST", {
        "namecharacter": nameCharacter,
        "ismale": sex
    }).then(response => {
        location.reload();


    }).catch(error => {
        errorCreateCharacter.style.color = "red"; 
        if (error.status == 409) {
            errorCreateCharacter.innerHTML = "Personaje ya existe en la BD"
        } else if (error.status == 428) {
            errorCreateCharacter.innerHTML = "ya tienes muchas cuentas creadas"
        } else{
               errorCreateCharacter.innerHTML = "Problemas interno del Servidor"
        }
    });
}