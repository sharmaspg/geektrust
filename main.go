package main

import (
	"fmt"
	"geektrust/util"
	"os"

	"strings"
)


func main(){
	var allies []string

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	for _,v := range util.ReadFile(os.Args[1]){
		data := strings.Split(v," ")
		if len(data)< 2 {
			fmt.Println(fmt.Errorf("format is not correct"))
			return
		}else {
			message := util.FormatMessage(v)
			decrypted := util.DecryptMessage(message)
			message.DecryptedMessage = decrypted
			if util.IsValid(message) && !util.Contains(allies,message.Kingdom) {
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
