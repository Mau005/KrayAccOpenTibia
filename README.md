# KrayAccOpenTibia – Creador de Cuentas para Servidores OpenTibia (TFS 1.6.x)

## Descripción del proyecto

**KrayAccOpenTibia** es una aplicación web escrita en Go que funciona como un **creador y administrador de cuentas (Account Creator)** para servidores OpenTibia basados en **The Forgotten Server 1.6.x**. Ofrece una interfaz web moderna para registrar, iniciar sesión y gestionar cuentas y personajes, y actúa como **servidor de login unificado** para múltiples mundos. **DEPRECADO**

## Características clave

- Soporte multi-mundo (sincronización entre varios servidores)
- Registro e inicio de sesión de cuentas
- Gestión de cuentas y personajes desde el navegador
- Generación automática de tokens de sesión (login para cliente Tibia 11/12)
- Manejo de cuenta Premium por tiempo
- Creación de personajes con parámetros configurables
- Sistema de noticias y noticias para cliente Tibia
- Highscores (rankings) por mundo y skill
- Estado del servidor y jugadores online
- Servidor de login (compatible con OTC/Tibia oficial)
- API REST para sincronización entre instancias
- Opción para lanzar y reiniciar el TFS automáticamente

## Requisitos

- Go (1.20+)
- MySQL/MariaDB
- Esquema de base de datos de TFS 1.6.x
- Configuración de TFS (config.lua)
- (Opcional) Certificado TLS
- (Opcional) Variable de entorno `KRAY_PASSWORD` para multi-mundo

## Instalación

```bash
git clone https://github.com/Mau005/KrayAccOpenTibia.git
cd KrayAccOpenTibia
go build -o KrayAccWeb .
./KrayAccWeb
```

Edita el archivo `config.yml` para definir:
- Conexión a la base de datos
- Mundo local (o sin mundo si es sólo API)
- Lista de mundos remotos (en `PoolServer`)
- Parámetros por defecto para nuevos personajes
- TLS y otras opciones

## Multi-Mundo

Para sincronización entre varios mundos:
1. Ejecutar una instancia por mundo (una principal con UI, las otras como API)
2. Configurar `KRAY_PASSWORD` en todas las instancias
3. En la principal, llenar `PoolServer` con IPs y tokens de los otros mundos
4. Al registrar cuentas o personajes, estos se replicarán automáticamente

## Uso

- Registro de cuentas desde la web
- Inicio de sesión con nombre o email
- Crear personajes eligiendo mundo
- Ver lista de personajes y VIP status
- Compatible con clientes Tibia 11/12 usando login server
- Rankings, noticias y online list

## Créditos

Desarrollado por **Mau005 / Krayno**  
Basado en el protocolo y estructura de TFS 1.6.x.  

---

¡Disfruta de KrayAccOpenTibia!  
Contribuciones y feedback son bienvenidos.
