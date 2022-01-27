package util

const (
	Gorilla ="Gorilla"
	Panda = "Panda"
	Octopus = "Octopus"
	Mammoth = "Mammoth"
	Owl ="Owl"
	Dragon = "Dragon"
	SPACE = "SPACE"
	LAND = "LAND"
	WATER = "WATER"
	ICE = "ICE"
	AIR = "AIR"
	FIRE = "FIRE"
)

func GetEmblem(kingdom string) string {
	switch kingdom {
	case SPACE: return Gorilla
	case LAND:return Panda
	case WATER:return Octopus
	case ICE:return Mammoth
	case AIR:return Owl
	case FIRE:return Dragon
	default:
		return ""
	}
}