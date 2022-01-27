package main

import (
	"github.com/SaravananPitchaimuthu/Practice/Chapter7/gadget"
)

func playList(player Player, mixTape []string) {
	for _, song := range mixTape {
		player.Play(song)
	}
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder) //type assertions
	if ok {
		recorder.Record()
	}
}

type Player interface {
	Play(string)
	Stop()
}

func main() {
	var player Player
	var Recorder Player
	player = gadget.TapePlayer{}
	mixTape := []string{"Oo soldriya mama Oo Oh Soldriya mama", "Neruppe siki mukki Neruppee", "Thee pidika Thee pidika"}
	playList(player, mixTape)
	Recorder = gadget.TapeRecorder{}
	playList(Recorder, mixTape)
}
