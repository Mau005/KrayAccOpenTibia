function highScore(target) {
    let world = document.getElementById("world").value;
    let skills = document.getElementById("skills").value;
    const tableBody = document.getElementById("tableHighScore").querySelector("tbody");
    tableBody.innerHTML = "";
    sendRequest("/highscore/"+world+"/"+skills, "POST", {
    }).then(response => {
        console.log(response)

        response.Players.forEach(element => {
            const row = document.createElement("tr");
            const nameCell = document.createElement("td");
            const skillsCell = document.createElement("td");
            nameCell.textContent = element.name;

            switch (skills) {
                case "0":
                skillsCell.textContent = element.skill_fist;
                break;
                case "1":
                skillsCell.textContent = element.skill_club;
                break;
                case "2":
                skillsCell.textContent = element.skill_axe;
                break;
                case "3":
                skillsCell.textContent = element.skill_sword;
                break;
                case "4":
                skillsCell.textContent = element.skill_dist;
                break;
                case "5":
                skillsCell.textContent = element.skill_shielding;
                break;
                case "6":
                skillsCell.textContent = element.skill_fishing;
                break;
                case "7":
                skillsCell.textContent = element.maglevel;
                break;
                case "8":
                skillsCell.textContent = element.level;
                break;
                default:
                    alert("entro en default");
                    break;
            }
            row.appendChild(nameCell);
            row.appendChild(skillsCell);
            tableBody.appendChild(row);
        });
    }).catch(error => {

        console.log(error)

    });
}

function getPlayerImageSource(urlTarget,player) {
    const baseUrl = urlTarget; // Assuming this is a global variable
    const urlParams = {
      id: player.looktype,
      addons: player.lookaddons,
      head: player.lookhead,
      body: player.lookbody,
      legs: player.looklegs,
      feet: player.lookfeet,   
  
      mount: 0, // Assuming mount is always 0
      direction: 3, // Assuming direction is always 3
    };
  
    const url = new URL(baseUrl + "/animoutfit.php", window.location.origin);
    url.search = new URLSearchParams(urlParams); // Add query parameters
  
    return `<img src="${url.toString()}" alt="Jugador ${player.ID}">`;
  }