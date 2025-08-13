import hashlib
from pathlib import Path
import json

def preparingHash(file_bytes: bytes) -> str:
    """
    Calcula el hash SHA-512 de los bytes del archivo.
    """
    m = hashlib.sha512()
    m.update(file_bytes)
    return m.hexdigest()

def preparingFile(file: Path) -> dict:
    """
    Prepara un diccionario con el nombre del archivo, su hash, y su tipo (archivo).
    """
    return {
        "Name": file.name,
        "Hash": preparingHash(file.read_bytes()),  # Corrige la llamada para calcular el hash
        "Type": True  # True para indicar que es un archivo
    }

def preparingDir(directory: Path) -> dict:
    """
    Prepara un diccionario con el nombre del directorio y los archivos dentro.
    """
    return {
        "Name": directory.name,
        "Type": False,  # False para indicar que es un directorio
        "Files": []
    }

def calculate_data(path: Path) -> dict:
    """
    Función recursiva para calcular el hash de los archivos y crear el JSON de control de actualizaciones.
    """
    if path.is_file():
        return preparingFile(path)
    
    # Es un directorio
    dir_info = preparingDir(path)
    for element in path.iterdir():
        if element.name in ["versiones.json", "main.py"]:
            continue  # Ignorar archivos o carpetas no deseadas
        
        # Llamada recursiva para archivos y subdirectorios
        dir_info["Files"].append(calculate_data(element))
    
    return dir_info

# Ruta inicial
p = Path(".")

# Generar la estructura de datos
project_structure = calculate_data(p)

# Guardar en un archivo JSON
with open('versiones.json', 'w') as json_file:
    json.dump(project_structure, json_file, indent=4)

print(f"JSON de control de actualizaciones generado en: {Path('versiones.json').absolute()}")
