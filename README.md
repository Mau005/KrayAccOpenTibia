# KrayAccOpenTibia – Account Creator for OpenTibia Servers (Canary)

## Project Description

**KrayAccOpenTibia** is a web application written in Go that functions as an **Account Creator** for OpenTibia servers based on **Canary Server**. It offers a modern web interface for registering, logging in, and managing accounts and characters, and acts as a **unified login server** for multiple worlds.

## Key Features

- Multi-world support (synchronization between multiple servers)
- Account registration and login
- Account and character management from the browser
- Automatic session token generation (login for Tibia 12/13/14 client)
- Timed Premium Account management
- Character creation with configurable parameters
- News and feed system for the Tibia client
- High scores (rankings) by world and skill
- Server status and online players
- Login server (compatible with OTC/official Tibia)
- REST API for synchronization between instances
- Option to automatically launch and restart Canary
- Added News Short
- Added manifest to update clients at the atomic level
- Added manifest tool to create manifest and host independently of the website

## Requirements

- Go (1.20+)
- MySQL/MariaDB
- Canary database schema 1.6.x
- Canary configuration (config.lua)
- (Optional) TLS certificate
- (Optional) `KRAY_PASSWORD` environment variable for multiworld

## Manifest
**ManifestTools** It is used with the launcher to be able to have a client automation session, this process is atomic level on the part of the launcher since it is configured only to download what is necessary

### Maniferst Config

Dev:
```bash
go run .\cmd\manifest\ -create_manifest
```
will create a manifest of all the data added to the .\client folder in the project root

Dev:
```bash
go run .\cmd\manifest\
```
It will run a mini server, independent of the web, for other OpenTibia projects or any type of program that requires it, obviously it will use the same configuration as the standard KrayACC project.

```json
{
  "app": "AinhoOT",
  "version": "1.0.0",
  "base_url": "https://ainho.ddns.net/launcher_client",
  "files": [
    {
      "path": "sounds/sounds-926b9436418eb757089bbc3600e889d60f7b211f065eaf308216676fb8e61807.dat",
      "size": 90882,
      "sha256": "926b9436418eb757089bbc3600e889d60f7b211f065eaf308216676fb8e61807"
    }
  ]
}
```
- app: "project name"
- version: "versioning to be used"
- base_url: "available static path where the client will search for the files"
- files: list of hashed files with a size to be verified by the client

## Installation
Easy Run all OS:
```bash
git clone https://github.com/Mau005/KrayAccOpenTibia.git
cd KrayAccOpenTibia
go run .\cmd\server\main.go
```
Example Build All OS:
-Build
```bash
go build .\cmd\server\main.go
```


Compile Proyect: 
On Windows:
```bash
.\build.ps1 -Target windows-amd64
```
On Linux/Mac:
```bash
chmod +x build.sh
./build.sh -t linux-amd64 -v 1.0.0
```

Edit the `config.yml` file to define:
- Database connection
- Local world (or no world if it's just an API)
- List of remote worlds (in `PoolServer`)
- Default parameters for new characters
- IP Static tunnel active
- TLS and other options


## Multiworld

To sync between multiple worlds:
1. Run one instance per world (one main instance with UI, the others as API)
2. Set `KRAY_PASSWORD` on all instances
3. On the main instance, populate `PoolServer` with IPs and tokens from the other worlds
4. When registering accounts or characters, these will be automatically replicated

## Usage

- Registering accounts from the web
- Logging in with name or email
- Creating characters by choosing a world
- Viewing the character list and VIP status
- Compatible with Tibia 11/12 clients using the login server
- Rankings, news, and online lists

## Credits

Developed by **Mau005 / Krayno**
Based on the Canary protocol.

---

Enjoy KrayAccOpenTibia!
Contributions and comments are welcome.
