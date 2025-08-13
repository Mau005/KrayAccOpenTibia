package components

import (
	"fmt"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/src/config"
	"github.com/Mau005/KrayAccOpenTibia/src/controller"
	"github.com/Mau005/KrayAccOpenTibia/src/models"
	"github.com/Mau005/KrayAccOpenTibia/src/utils"
)

func CreateLastPlayerKills() string {
	var playerCtl controller.PlayerController
	itemPlayers := ""
	players := playerCtl.GetPlayerDeath()

	for index, death := range players {
		time := time.Unix(int64(death.Time), 0)

		itemPlayers += fmt.Sprintf(`
		
										<tr>
                                            <td>%d</td>
                                            <td>%s</td>
                                            <td>%s</td>
                                        </tr>
		`, index+1,
			fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second()),
			fmt.Sprintf("<a href='/get_character/%s'>%s</a> died at level %d by a %s.",
				death.Player.Name,
				death.Player.Name,
				death.Level,
				death.KilledBy))
	}

	return fmt.Sprintf(`
	
		<h1> Ultimas Muertes!?</h1>
                        <!-- Tabla de Lista de Personajes -->
                        <div class="card mb-4">
                            <div class="card-header">
                                Ultimas Muertes!
                            </div>
                            <div class="card-body">
                                <table class="table">
                                    <thead>
                                        <tr>
                                            <th scope="col">N°</th>
                                            <th scope="col">Fecha</th>
                                            <th scope="col">Descripción</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        %s
                                    </tbody>
                                </table>
                            </div>
                        </div>
	`, itemPlayers)

}

func CreatePlayerOnline() string {
	var PoolConnectionCtl controller.PoolConnectionController

	whoIsOnline := `
    <h1>Quien esta Online?</h1>
    %s
    `

	data := PoolConnectionCtl.WhoIsOnlinePoolConnection()
	contentOutput := ""
	for world, players := range data {
		content := `
        
                        <!-- Acordeón -->
                        <div class="accordion mb-4">
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="headingOne">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#%s" aria-expanded="true" aria-controls="%s">
                                        %s
                                    </button>
                                </h2>
                                <div id="%s" class="accordion-collapse collapse" aria-labelledby="headingOne" data-bs-parent="#%s">
                                    <div class="accordion-body">
                                        %s
                                    </div>
                                </div>
                            </div>
                        </div>
        `
		playerList := ""

		for _, player := range players {
			playerList += fmt.Sprintf(`
		
            <tr>
                <td>%s</td>
                <td>%s</td>
                <td>%d</td>
                <td>%d</td>
            </tr>
            `, FunctionImagenSourcePlayer(player), player.Name, player.Level, player.Experience)
		}
		contentOutput += fmt.Sprintf(content, world, world, world, world, world, fmt.Sprintf(`
                        <!-- Tabla de Lista de Personajes -->
                        <div class="card mb-4">
                            <div class="card-header">
                                Lista de Personajes
                            </div>
                            <div class="card-body">
                                <table class="table">
                                    <thead>
                                        <tr>
                                            <th scope="col">Outfits</th>
                                            <th scope="col">Nombre</th>
                                            <th scope="col">Nivel</th>
                                            <th scope="col">Experiencia</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        %s
                                        
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    `, playerList))
	}

	return fmt.Sprintf(whoIsOnline, contentOutput)

}

func CreateTopPlayerComponent(countPlayer int) string {
	var playerCtl controller.PlayerController
	players := playerCtl.GetPlayerLimits(countPlayer)

	liItems := ""

	for value, player := range players {
		urlImage := fmt.Sprintf("%s/animoutfit.php?id=%d&addons=%d&head=%d&body=%d&legs=%d&feet=%d&mount=0&direction=3",
			config.Global.ServerWeb.UrlOutfitsView, player.LookType, player.LookAddons, player.LookHead, player.LookBody, player.LookLegs, player.LookFeet)
		liItems += fmt.Sprintf(`
<li class="list-group-item">
	<img src="%s" alt="Jugador %d"> %s - Nivel %d
</li>
		`, urlImage, value, player.Name, player.Level)
	}

	componentes := `
<h4>Top %d Jugadores</h4>
<ul class="list-group">
	%s
</ul>
	
	`

	return fmt.Sprintf(componentes, countPlayer, liItems)
}

func CreateMyPlayers(navWeb models.NavWeb) string {
	itemPlayers := ""
	components := ""
	for _, player := range navWeb.MyPlayers {
		urlImage := fmt.Sprintf("%s/animoutfit.php?id=%d&addons=%d&head=%d&body=%d&legs=%d&feet=%d&mount=0&direction=3",
			config.Global.ServerWeb.UrlOutfitsView, player.LookType, player.LookAddons, player.LookHead, player.LookBody, player.LookLegs, player.LookFeet)
		itemPlayers += fmt.Sprintf(`
							<li class="list-group-item">
                                %s
                                %s - Nivel %d
                            </li>
		`, urlImage, player.Name, player.Level)

		components += itemPlayers
	}

	return fmt.Sprintf(`
	 					<ul class="list-group">
                            %s
                        </ul>
	`, components)
}

func CreateGetPlayer(player models.Players) string {

	sex := "Male"
	if player.Sex == 0 {
		sex = "Female"
	}

	playerStatic := fmt.Sprintf(`
    <div class="col-md-6">
      <div class="info-block">
        <table class="table table-bordered">
          <tbody>
            <tr><th>Nombre:</th><td style="color: limegreen;">%s</td></tr>
            <tr><th>Sexo:</th><td>%s</td></tr>
            <tr><th>Profession:</th><td>%s</td></tr>
            <tr><th>Level:</th><td>%d</td></tr>
            <tr><th>Residence:</th><td>%s</td></tr>
            <!--<tr><th>Guild:</th><td>Member of the <a href="#">Fuck This</a></td></tr>-->
            <tr><th>Last login:</th><td>%s CEST</td></tr>
            <!-- <tr><th>Created:</th><td>Jan 01 1970, 00:00:00 CEST</td></tr> -->
            <tr><th>Account:</th><td>Nothing</td></tr>
            <tr><th>World:</th><td>%s</td></tr>
          </tbody>
        </table>
      </div>
    </div>
    `, player.Name,
		sex,
		utils.FunctionGetVocation(player),
		player.Level,
		utils.GetTownGeneral(player.TownID),
		time.Unix(int64(player.LastLogin), 0).Format("Jan 02 2006, 15:04:05 MST"),
		player.World,
	)
	rlImage := fmt.Sprintf("%s/animoutfit.php?id=%d&addons=%d&head=%d&body=%d&legs=%d&feet=%d&mount=0&direction=3",
		config.Global.ServerWeb.UrlOutfitsView, player.LookType, player.LookAddons, player.LookHead, player.LookBody, player.LookLegs, player.LookFeet)

	imgPlayer := fmt.Sprintf(`
    <div class="col-md-6 text-center">
      <div class="outfit-frame p-3 border rounded">
        <img src="%s"
          alt="Animated Outfit"
          style="width: 300px; height: auto; image-rendering: pixelated;"
          class="img-fluid rounded">
      </div>
    </div>
    `, rlImage)

	skills := fmt.Sprintf(`
    <div class="col-md-4">
      <div class="info-block">
        <h4 style="color: var(--highlight-color);">Skills</h4>
        <table class="table table-bordered">
          <tbody>
            <tr><th>Magic Level</th><td>%d</td></tr>
            <tr><th></th><td></td></tr>
            <tr><th>Fist Fighting</th><td>%d</td></tr>
            <tr><th>Club Fighting</th><td>%d</td></tr>
            <tr><th>Sword Fighting</th><td>%d</td></tr>
            <tr><th>Axe Fighting</th><td>%d</td></tr>
            <tr><th>Distance</th><td>%d</td></tr>
            <tr><th>Shielding</th><td>%d</td></tr>
            <tr><th>Fishing</th><td>%d</td></tr>
          </tbody>
        </table>
      </div>
    </div>
    `, player.MagLevel, player.SkillFist, player.SkillClub, player.SkillSword, player.SkillAxe, player.SkillDist, player.SkillShielding, player.SkillFishing)
	target := config.Global.ServerWeb.UrlItemView

	itemsEquipmen := utils.GetEquipmenItem(player.PlayerItems)
	equipment := fmt.Sprintf(`
<div class="col-md-8">
  <div class="info-block text-center">
    <h4 style="color: var(--highlight-color);">Equipment</h4>
    <div class="equipment-grid mx-auto">

      <div class="slot amulet"><img src="%s/%d.png" alt="Amulet"></div>
      <div class="slot helmet"><img src="%s/%d.png" alt="Helmet"></div>
      <div class="slot backpack"><img src="%s/%d.png" alt="Backpack"></div>

      <div class="slot weapon"><img src="%s/%d.png" alt="Weapon"></div>
      <div class="slot armor"><img src="%s/%d.png" alt="Armor"></div>
      <div class="slot shield"><img src="%s/%d.png" alt="Shield"></div>

      <div class="slot ring"><img src="%s/%d.png" alt="Ring"></div>
      <div class="slot legs"><img src="%s/%d.png" alt="Legs"></div>
      <div class="slot ammo"><img src="%s/%d.png" alt="Ammo"></div>
      <br>
      <div class="slot feet"><img src="%s/%d.png" alt="Feet"></div>
      <br>

    </div>
  </div>
</div>
`,
		target, itemsEquipmen[utils.AmuletEquipment],
		target, itemsEquipmen[utils.HeadEquipment],
		target, itemsEquipmen[utils.BackpackEquipment],

		target, itemsEquipmen[utils.LeftHandEquipment],
		target, itemsEquipmen[utils.ArmorEquipment],
		target, itemsEquipmen[utils.RightHandEquipment],

		target, itemsEquipmen[utils.RingEquipment],
		target, itemsEquipmen[utils.LegsEquipment],
		target, itemsEquipmen[utils.AmmoEquipment],

		target, itemsEquipmen[utils.FeetEquipment],
	)

	deathList := ""

	for _, death := range player.PlayerDeaths {
		time := time.Unix(int64(death.Time), 0)

		targetKill := death.KilledBy

		if death.IsPlayer {
			targetKill = fmt.Sprintf(`<a href="/get_character">%s</a>`, death.KilledBy)
		}
		deathList += fmt.Sprintf(`
    <tr>
      <td>%s</td>
      <td>%s</td>
      <td>Killed at level %d by a <strong>%s</strong>.</td>
    </tr>
    `,
			fmt.Sprintf("%d-%d-%d", time.Year(), time.Month(), time.Day()),
			fmt.Sprintf("%d:%d:%d", time.Hour(), time.Minute(), time.Second()),
			death.Level,
			targetKill,
		)
	}

	return fmt.Sprintf(`
<div class="container my-5 character-profile">
  <h2 class="text-center mb-4" style="color: var(--highlight-color);">Character Information</h2>

  <div class="row g-4">
    <!-- Info + outfit -->
    %s

    <!-- Outfit -->
    %s
  </div>

  <hr class="my-4" style="border-color: var(--highlight-color);">

  <div class="row g-4">
    <!-- Skills -->
    %s

    <!-- Equipment grid -->
    %s
  </div>

  <hr class="my-4" style="border-color: var(--highlight-color);">

  <!-- Deaths -->
  <div class="info-block">
    <h4 class="mb-3" style="color: var(--highlight-color);">Character Deaths</h4>
    <table class="table table-bordered">
      <thead>
        <tr>
          <th>Fecha</th>
          <th>Hora</th>
          <th>Descripción</th>
        </tr>
      </thead>
      <tbody>
      %s
      </tbody>
    </table>
  </div>
</div>

    `, playerStatic, imgPlayer, skills, equipment, deathList)
}
