package main

import (
	"fmt"
	"geektrust/util"

	//"geektrust/util"
	"io/ioutil"
	"os"
	"strings"
)

type Message struct {
	Kingdom string
	SecretMessage string
	DecryptedMessage string

}
func readFile() []string{

	// data is the file content, you can use it
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return nil
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	return strings.Split(string(data),"\n")
}

func formatMessage(message string) Message{
	kingdomMessage := strings.Split(message, " ")
	if len(kingdomMessage) < 2 {
		fmt.Println("Format is not correct")
	}
	return Message{
		Kingdom: kingdomMessage[0],
		SecretMessage: strings.TrimSpace(kingdomMessage[1]),
	}
}

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func decryptMessage(message Message) string{
	emblem := util.GetEmblem(message.Kingdom)
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

func isValid(message Message) bool{
	emblem := strings.ToUpper(util.GetEmblem(message.Kingdom))
	for _,rune := range message.DecryptedMessage {
		emblem = strings.Replace(emblem,strings.ToUpper(string(rune)),"",1)
	}
	if len(emblem)>0 {
		return false
	}
	return true
}

func main(){

	var allies []string
	for _,v := range readFile(){
		data := strings.Split(v," ")
		if len(data)< 2 {
			fmt.Println(fmt.Errorf("format is not correct"))
			return
		}else {
			message := formatMessage(v)
			decrypted := decryptMessage(message)
			message.DecryptedMessage = decrypted
			if isValid(message){
				allies = append(allies, message.Kingdom)
			}

		}
	}
	if len(allies) >=3 {
		fmt.Println(util.SPACE, strings.Join(allies," "))
	}else {
		fmt.Println("NONE")
	}

}
