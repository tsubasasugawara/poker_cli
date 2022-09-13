package terminal

import (
	"syscall"
	"log"
)

const (
	ENTER 	byte = 13
	CtrlC 	byte = 3

	C		byte = 99
	A 		byte = 97
	F		byte = 102

	ONE		byte = 49
	NINE	byte = 57
)

// バッファの値を取得する
func readBuffer(bufCh chan []byte) {
	buf := make([]byte, 1024)
	for {
		if n, err := syscall.Read(syscall.Stdin, buf); err == nil {
			bufCh <- buf[:n]
		}
	}
}

// 文字列の入力を受け取る
func GetChar() {
	bufCh := make(chan []byte, 1)
	go readBuffer(bufCh)

	running := true
	for running {
		p := <-bufCh

		switch p[0] {
		case CtrlC:
			running = false
		default:
			log.Println(p)
		}
	}
}
