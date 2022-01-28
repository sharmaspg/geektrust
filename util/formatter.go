package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func ReadFile(filename string) []string{
	// data is the file content, you can use it
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		return nil
	}
	return strings.Split(string(data),"\n")
}


type Message struct {
	Kingdom string
	SecretMessage string
	DecryptedMessage string

}

func FormatMessage(message string) Message{
	kingdomMessage := strings.SplitN(message, " ",2)
	return Message{
		Kingdom: kingdomMessage[0],
		SecretMessage: strings.ReplaceAll(kingdomMessage[1]," ",""),
	}
}


func Contains(original []string, value string) bool {
	for _,v := range original {
		if v == value{
			return true
		}
	}
	return false
}