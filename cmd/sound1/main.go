package main

import (
	"github.com/krsanky/go-sound1"
)

func main() {
	err := sound1.Run1()
	if err != nil {
		panic(err)
	}
}
