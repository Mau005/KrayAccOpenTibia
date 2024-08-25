function loginUser() {
    let user = document.getElementById("username").value;
    let password = document.getElementById("password").value;
    const errorLogin = document.getElementById("errorLogin")
    if (password == "" || user == "") {
        document.getElementById("spanlogin").innerHTML = "Credenciales incorrectas"
    }

    sendRequest("/login", "POST", {
        "user": user,
        "password": password
    }).then(response => {
        location.reload();


    }).catch(error => {
        errorLogin.style.color = "red"; // Cambia el color a rojo
        if (error.status == 409) {
            errorLogin.innerHTML = "Credenciales incorrectas"
        } else {
            errorLogin.innerHTML = "Problemas interno del Servidor"
        }

        console.log(error)

    });
}