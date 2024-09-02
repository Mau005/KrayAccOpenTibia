package components

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

func CreateNewsComponents(navWeb models.NavWeb, limitedTicket int) string {
	//
	componentsTickets := ""
	// if navWeb.Authentication {
	// 	editOrName = ""
	// }else {

	// }
	var newsTicketCrl controller.NewsTickerController

	newsTicket := newsTicketCrl.GetTickerLimited(limitedTicket)
	for _, ticket := range newsTicket {
		editOrName := fmt.Sprintf("<h6>By %s</h6>", ticket.Player.Name)
		if navWeb.TypeAccess >= 5 {
			editOrName = `<button class="btn btn-primary">Delete</button>`
		}
		componentsTickets += fmt.Sprintf(`
                        <div class="news-item">
                            <img src="www/img/%s" alt="Imagen de Noticia">
                            <p>%s</p>
                            %s
                        </div>	
		
		`, newsTicketCrl.GetIconID(ticket.IconID), ticket.Ticket, editOrName)
	}

	return fmt.Sprintf(`
	                    <h1>Noticias</h1>
                        <!-- Componente de Noticias Previas -->
						%s
	`, componentsTickets)
}
