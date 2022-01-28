package util

import "testing"

func TestDecryptMessage(t *testing.T) {
	message := Message{
		Kingdom: LAND,
		SecretMessage: "FGHIJK",
	}
	wanted := "ABCDEF"
	recieved := DecryptMessage(message)
	if recieved != wanted {
		t.Errorf("got %q, wanted  %q",recieved, wanted)
	}
}

func TestDecryptMessage2(t *testing.T) {
	message := Message{
		Kingdom: LAND,
		SecretMessage: "ABCDEF",
	}
	wanted := "VWXYZA"
	recieved := DecryptMessage(message)
	if recieved != wanted {
		t.Errorf("got %q, wanted  %q",recieved, wanted)
	}
}

func TestIsValid(t *testing.T) {
	message := Message{
		Kingdom: AIR,
		SecretMessage: "ROZO",
	}
	msg := DecryptMessage(message)
	message.DecryptedMessage = msg
	if !IsValid(message) {
		t.Errorf("got false, wanted  true")
	}
}


func TestIsValid2(t *testing.T) {
	message := Message{
		Kingdom: AIR,
		SecretMessage: "OWLAOWLBOWLC",
	}
	msg := DecryptMessage(message)
	message.DecryptedMessage = msg
	if IsValid(message) {
		t.Errorf("got true, wanted  false")
	}
}