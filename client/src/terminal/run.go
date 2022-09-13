package terminal

import (
	"log"

	"github.com/nsf/termbox-go"
)

func Run() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	GetChar()
}
