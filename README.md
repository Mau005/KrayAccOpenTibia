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

## Requirements

- Go (1.20+)
- MySQL/MariaDB
- Canary database schema 1.6.x
- Canary configuration (config.lua)
- (Optional) TLS certificate
- (Optional) `KRAY_PASSWORD` environment variable for multiworld

## Installation

```git clone https://github.com/Mau005/KrayAccOpenTibia.git
cd KrayAccOpenTibia
go to build -o KrayAccWeb
./KrayAccWeb
```

Edit the `config.yml` file to define:
- Database connection
- Local world (or no world if it's just an API)
- List of remote worlds (in `PoolServer`)
- Default parameters for new characters
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
Based on the Canary protocol and framework.

---

Enjoy KrayAccOpenTibia!
Contributions and comments are welcome.
