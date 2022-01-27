package gadget

import "fmt"

type TapeRecorder struct {
	Microphones int
}

type TapePlayer struct {
	Batteries string
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing the song", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopping")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing the song", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopping")
}
