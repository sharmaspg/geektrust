package util

import "testing"

func TestGetEmblem(t *testing.T) {
	data := map[string]string{
		SPACE:Gorilla,
		ICE:Mammoth,
		LAND:Panda,
		WATER:Octopus,
		AIR:Owl,
		FIRE:Dragon,
	}
	for kingdom,emblem := range data {
		emb := GetEmblem(kingdom)
		if emb != emblem {
			t.Errorf("got %q,wanted %q", emb, emblem)
		}
	}
}

func TestGetEmblem2(t *testing.T) {
	if GetEmblem("") != ""{
		t.Errorf("got %q, wanted none",GetEmblem(""))
	}
}

func TestContains(t *testing.T) {
	if !Contains([]string{"A","B","C","D"},"A") {
		t.Errorf("Wanted true, got false")
	}
}

func TestContains2(t *testing.T) {
	if Contains([]string{"TestCase"},"Not in Slice"){
		t.Errorf("wanted false, got true")
	}
}


func TestFormatMessage(t *testing.T) {
	Sampletext := "KINGDOM Message 2465"
	data := FormatMessage(Sampletext)
	if data.SecretMessage != "Message2465"{
		t.Errorf("Wanted Message2465, got %q",data.SecretMessage)
	}

}

func TestReadFile(t *testing.T) {
	received := ReadFile("Notfilename.txt")
	if received != nil {
		t.Errorf("got %q, wanted nil", received)
	}
}

func TestReadFile2(t *testing.T) {
	received := ReadFile("sample.txt")
	if received == nil {
		t.Errorf("wanted []string, got nil")
	}
}
