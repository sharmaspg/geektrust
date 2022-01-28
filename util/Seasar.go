package util

import "strings"

func DecryptMessage(message Message) string{
	emblem := GetEmblem(message.Kingdom)
	decryptedMessage := ""
	for _,a := range message.SecretMessage{
		alphabetIndex := strings.Index(Alphabet,string(a))
		newIndex := alphabetIndex - len(emblem)
		if newIndex >=0 {
			decryptedMessage = decryptedMessage + string(Alphabet[newIndex])
		}else {
			decryptedMessage = decryptedMessage + string(Alphabet[len(Alphabet)+newIndex])
		}

	}

	return decryptedMessage
}

func IsValid(message Message) bool{
	emblem := strings.ToUpper(GetEmblem(message.Kingdom))
	for _,rune := range message.DecryptedMessage {
		emblem = strings.Replace(emblem,strings.ToUpper(string(rune)),"",1)
	}
	if len(emblem)>0 {
		return false
	}
	return true
}
