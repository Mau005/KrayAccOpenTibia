function validateInsults(username) {
    const forbiddenWords = [
        "pene", "wn", "puta", "puto", "cabron", "mierda", "imbecil", "idiota", "culiao",
        "conchetumare", "maricon", "huevon", "perra", "zorra", "culero", "joto",
        "csm", "hdp", "ctm", "tonto", "estupido", "pendejo", "tarado",
        "fuck", "shit", "bitch", "asshole", "bastard", "dick", "cunt", "faggot",
        "slut", "whore", "damn", "moron", "idiot", "retard",
        "puta", "bosta", "merda", "caralho", "cacete", "piranha", "viado",
        "desgraca", "imbecil", "idiota", "otario", "babaca",
        "Community", "Manager", "CM", "GM", "GOD", "gamemaster"
    ];
    
    const forbiddenWordsRegex = new RegExp(`\\b(${forbiddenWords.join("|")})\\b`, "i");

    return !forbiddenWordsRegex.test(username);
}

// Función para validar caracteres especiales y formato
// Función para validar caracteres especiales, acentos y formato
function validateSpecialCharsAndSpaces(username) {
    // Verificar caracteres especiales y acentos
    const specialCharsRegex = /[^a-zA-Z0-9 ]/;
    // Verificar más de dos espacios consecutivos
    const multipleSpacesRegex = / {3,}/;
    // Verificar si comienza con un espacio
    const leadingSpaceRegex = /^ /;

    // Si hay caracteres especiales, acentos, más de dos espacios consecutivos o el nombre empieza con un espacio, es inválido
    if (specialCharsRegex.test(username) || multipleSpacesRegex.test(username) || leadingSpaceRegex.test(username)) {
        return false;
    }

    return true;
}

function createCharacter() {
    let nameCharacter = document.getElementById("nameCharacter").value;
    let isMale = document.getElementById("radioMale").checked;
    let isFemale = document.getElementById("radioFemale").checked;
    const errorCreateCharacter = document.getElementById("errorCreateCharacter");
    errorCreateCharacter.innerHTML = "";
    errorCreateCharacter.style.color = "red"; // Cambia el color a rojo

    if (validateInsults(nameCharacter) == false){
        errorCreateCharacter.innerHTML = "Nombre no puede tener insultos o palabras claves por el staff"
        return
    }

    if (validateSpecialCharsAndSpaces(nameCharacter) == false){
        errorCreateCharacter.innerHTML = "Nombre no puede tener caracteres especiales o espacios innesesarior"
        return
    }


    if (isMale == false && isFemale == false) {
        const errorElement = document.getElementById("errorCreateCharacter");
        errorCreateCharacter.innerHTML = "Tiene que indicar que sexo tiene";
        
        return
    }
    let sex = 0; 
    if (isMale) {
        sex = 1;
    }
    alert("paso lo filtros");
    return
    sendRequest("/auth/create_character", "POST", {
        "namecharacter": nameCharacter,
        "ismale": sex
    }).then(response => {
        location.reload();


    }).catch(error => {
        if (error.status == 409) {
            errorCreateCharacter.innerHTML = "Personaje ya existe en la BD"
        } else if (error.status == 428) {
            errorCreateCharacter.innerHTML = "ya tienes muchas cuentas creadas"
        } else{
               errorCreateCharacter.innerHTML = "Problemas interno del Servidor"
        }
    });
}