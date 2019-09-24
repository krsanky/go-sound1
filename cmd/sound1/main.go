package main

import (
	"oldcode.org/home/wise/repo/go/sound1"
)

func main() {
	err := sound1.Run1()
	if err != nil {
		panic(err)
	}
}
