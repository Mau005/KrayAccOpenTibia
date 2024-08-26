let timerInterval;
function startCounter(initialSeconds) {
    if (initialSeconds == ""){
        document.getElementById("counter").innerText = "Server OFF"
        return
    }
    
    let seconds = parseInt(initialSeconds);

    if (timerInterval) {
        clearInterval(timerInterval);
    }

    // Función para incrementar el tiempo
    function incrementTime() {
        let hours = Math.floor(seconds / 3600);
        let remainingSeconds = seconds % 3600;
        let minutes = Math.floor(remainingSeconds / 60);
        let finalSeconds = remainingSeconds % 60;

        let displayText = "";

        if (hours > 0) {
            displayText += `${hours} Horas`;
        }

        if (minutes > 0) {
            if (displayText) displayText += ", ";
            displayText += `${minutes} Minutos`;
        }

        if (finalSeconds > 0 || seconds === 0) {
            if (displayText) displayText += " y ";
            displayText += `${finalSeconds} Segundos`;
        }

        document.getElementById("counter").innerText = displayText;

        seconds++;
    }

    // Iniciar el contador
    incrementTime();
    timerInterval = setInterval(incrementTime, 1000);
}