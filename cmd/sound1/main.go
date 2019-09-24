package main

import sound1 "github.com/krsanky/go-sound1"

func main() {
	err := sound1.Run1()
	if err != nil {
		panic(err)
	}
}
