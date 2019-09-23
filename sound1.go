package sound1

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
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

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

}
