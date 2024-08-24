function sendRequest(url, method, data) {
    return new Promise((resolve, reject) => {
        const options = {
            method: method,
            headers: {
                'Content-Type': 'application/json',
                // Puedes agregar más encabezados según sea necesario
            },
        };

        if (data) {
            options.body = JSON.stringify(data);
        }

        fetch(url, options)
            .then(response => {
                // Verifica si la respuesta indica éxito (código de estado 2xx)
                if (response.ok) {
                    // Parsea la respuesta JSON y resuelve la promesa
                    return response.json().then(resolve);
                } else {
                    // Rechaza la promesa con un objeto de error que contiene el código de estado
                    reject({ status: response.status, statusText: response.statusText });
                }
            })
            .catch(error => {
                // Rechaza la promesa con el error de red
                reject (error);
            });
    });
}