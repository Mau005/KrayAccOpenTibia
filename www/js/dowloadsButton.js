function downloadFile(url, nameFile) {
// Crea un enlace temporal
    const link = document.createElement('a');
    link.href = url;
    link.download = nameFile;

    // Añade el enlace al documento
    document.body.appendChild(link);

    // Hace clic en el enlace para iniciar la descarga
    link.click();

    // Remueve el enlace temporal
    document.body.removeChild(link);
}