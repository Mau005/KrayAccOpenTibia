function registerAccount() {
    let user = document.getElementById("regUsername").value;
    let email = document.getElementById("email").value;
    let password = document.getElementById("regPassword").value;
    let passwordtwo = document.getElementById("confirmPassword").value;
    let terms = document.getElementById("terms").checked;

    const errorSpan = document.getElementById("errorRegAccount")
    errorSpan.style.color = "red"; // Cambia el color a rojo

    if (password.length  <= 5){
        errorSpan.innerHTML = "Credenciales muy corta"
        return
    }
    if (user.length <= 3){
        errorSpan.innerHTML = "Nombre muy corto"
        return
    }
    if (password != passwordtwo) {
        errorSpan.innerHTML = "Credenciales no son iguales"
        return
    }
    if (terms == false){
        errorSpan.innerHTML = "sebe aceptar nuestras politicas"
        return
    }

    sendRequest("/create_account", "POST", {
        "IsTerms": user,
        "Password": password,
        "IsTerms": terms,
        "PasswordTwo": passwordtwo,
        "Email": email
    }).then(response => {
        location.reload();
    }).catch(error => {
        if (error.status == 406) {
            errorSpan.innerHTML = "Usuario o correo electronicos ya usados"
        } else {
            errorSpan.innerHTML = "Problemas interno del Servidor"
        }

        console.log(error)

    });
}