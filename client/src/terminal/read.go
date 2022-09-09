package terminal

import (
	"syscall"
	
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
		if p != "" {
			log.Println(p)
		}
	}
}