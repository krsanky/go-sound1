package sound1

import (
	"fmt"
	"os"

	"github.com/faiface/beep/mp3"
)

func Noise() {
	//chair-on-fire.mp3
	f, err := os.Open("data/awright.mp3")
	if err != nil {
		panic(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()
	fmt.Printf("streamer:%v format:%v\n", streamer, format)
}
