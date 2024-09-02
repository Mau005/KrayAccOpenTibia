package handler

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/components"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type Layouthandler struct{}

func (lh *Layouthandler) Generatelayout(navWeb models.NavWeb, condition models.SolicitudeLayout) (layout models.Layout) {

	//Authentication User register or create character
	if navWeb.Authentication {
		layout.Modal += components.CreateModalCreateCharacter()
		layout.Scripts += `
		<script src="/www/js/create_character.js"></script>`
	} else {
		layout.Modal += components.CreateModalRegister()
		layout.Scripts += `
		<script src="/www/js/register_account.js"></script>
		`
	}

	if condition.News {
		layout.News = components.CreateNewsComponents(navWeb, utils.LimitRecordFive)
	}
	if condition.Login {
		layout.Login = components.CreateLogin(navWeb)
		layout.Scripts += `
        <script src="/www/js/login_home.js"></script>
        <script src="/www/js/redirectMenuLogin.js"></script>
		`
	}
	if condition.ServerStatus {
		layout.ServerStatus = components.CreateServerStatus(controller.TempData.ServerStatus)
		layout.Scripts += fmt.Sprintf(`
        <script src="/www/js/uptime.js"></script>
        <script>
            startCounter('%s')
        </script>
		`, controller.TempData.ServerStatus.ServerInfo.Uptime)
	}

	if condition.WhoIsOnline {
		layout.WhoIsOnline = components.CreatePlayerOnline()
	}
	if condition.LastDeath {
		layout.LastDeath = components.CreateLastPlayerKills()
	}

	if condition.TopPlayers {
		layout.TopPlayers = components.CreateTopPlayerComponent(utils.LimitRecordFive)
	}
	if condition.Rates {
		layout.Rates = components.CreateRates(controller.TempData.ServerStatus)
	}

	//default:
	layout.NavBar = components.CreateNavbar(navWeb)
	layout.LogoButtons = `
			<div class="logo-container">
                <img src="www/img/logo.png" alt="Logo" class="logo">
            </div>
            <button onclick="downloadFile('/www/dowloads/TheLastRookgardBETA.zip', 'TheLastRookgardV1.0.0BETA.zip')" class="vibrant-button">Descargar</button>`
	layout.Head = `
	    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.1, user-scalable=no" />
        <title>The Last Rookgard By Krayno</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
        <link rel="stylesheet" href="/www/style.css">
        <link rel="icon" type="image/png" sizes="16x16" href="/www/img/icon.png">
    </head>`

	//Scripts default:
	layout.Scripts += `
        <script src="/www/js/promise.js"></script>
        <script src="/www/js/dowloadsButton.js"></script>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
	`

	return
}
