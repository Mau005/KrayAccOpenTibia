package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/Mau005/KrayAccOpenTibia/src/models"
	"github.com/fatih/color"
)

var ErrorColor = color.New(color.FgRed)

func Info(msg ...string) {

	infoColor := color.New(color.FgGreen).SprintFunc()
	fmt.Println(infoColor(uniteText("[OK]", msg)))
}

func Warn(msg ...string) {
	warnColor := color.New(color.FgHiYellow).SprintFunc()
	fmt.Println(warnColor(uniteText("[WARNING]", msg)))
}

func WarnLog(msg ...string) {
	warnColor := color.New(color.FgHiYellow).SprintFunc()
	fmt.Println(warnColor(uniteText("[ALERT SECURITY]", msg)))
}

func WarnSecurity(msg ...string) {
	warColor := color.New(color.BgYellow).SprintFunc()
	fmt.Println(warColor(uniteText("[WARNING-SECURITY]", msg)))
}

func ErrorR(msg ...string) {
	errorColor := color.New(color.FgRed).SprintFunc()
	fmt.Println(errorColor(uniteText("[ERROR]", msg)))
}
func ErrorFatal(msg ...string) {
	errorColor := color.New(color.FgRed).SprintFunc()
	log.Fatalln(errorColor(uniteText("[FATAL]", msg)))
}

func InfoBlue(msg ...string) {
	errorColor := color.New(color.FgBlue).SprintFunc()
	fmt.Println(errorColor(msg))
}

func InfoSuccess(msg string) {
	errorColor := color.New(color.FgHiCyan).SprintFunc()
	fmt.Println(errorColor(msg))
}
func InfoBlueNotLog(msg ...string) {
	errorColor := color.New(color.FgBlue).SprintFunc()
	fmt.Println(errorColor(msg))
}

func uniteText(target string, msg []string) string {
	msgComplex := []string{target}

	// Agregar los elementos de msg a msgComplex
	msgComplex = append(msgComplex, msg...)

	// Unir todos los elementos en una sola cadena con espacios entre ellos
	return strings.Join(msgComplex, " ")
}

func FunctionGetVocation(player models.Players) string {

	switch player.Vocation {
	case 0:
		return "No Vocation"
	case 1:
		return "Sorcerer"
	case 2:
		return "Druid"
	case 3:
		return "Paladin"
	case 4:
		return "Knight"
	case 5:
		return "Master Sorcerer"
	case 6:
		return "Elder Druid"
	case 7:
		return "Royal Paladin"
	case 8:
		return "Elite Knight"
	default:
		return "No encontrado"

	}
}

func GetTownGeneral(id int) string {

	switch id {
	case 1:
		return "Venore"
	case 2:
		return "Thais"
	case 3:
		return "Kazordoon"
	case 4:
		return "Carlin"
	case 5:
		return "Ab Dendriel"
	case 6:
		return "Rookgard"
	case 7:
		return "Liberty Bay"
	case 8:
		return "Port Hope"
	case 9:
		return "Ankrahmun"
	case 10:
		return "Darashia"
	case 11:
		return "Edron"
	case 12:
		return "Svargrond"
	case 13:
		return "Yalahar"
	case 14:
		return "Farmine"
	default:
		return "Not Defined"
	}
}

func GetEquipmenItem(items []models.PlayerItem) map[TypeEquipmentSlot]int {
	itemResult := make(map[TypeEquipmentSlot]int)
	for _, it := range items {
		if it.PID > 10 {
			continue
		}

		idItem := it.ItemType
		switch it.PID {
		case 1:
			itemResult[HeadEquipment] = idItem
		case 2:
			itemResult[AmuletEquipment] = idItem
		case 3:
			itemResult[BackpackEquipment] = idItem
		case 4:
			itemResult[ArmorEquipment] = idItem
		case 5:
			itemResult[RightHandEquipment] = idItem
		case 6:
			itemResult[LeftHandEquipment] = idItem
		case 7:
			itemResult[LegsEquipment] = idItem
		case 8:
			itemResult[FeetEquipment] = idItem
		case 9:
			itemResult[RingEquipment] = idItem
		case 10:
			itemResult[AmmoEquipment] = idItem
		}
	}
	return itemResult
}
