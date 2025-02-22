package main

import (
	"maths/maths"
	"os"
)

func main() {
	err := maths.HandleArgs(os.Args)
	if err != nil {
		return
	}
}
