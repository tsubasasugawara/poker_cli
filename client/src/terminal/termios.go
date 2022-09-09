package terminal

import (
	"syscall"

	"github.com/pkg/term/termios"
)

// 起動時のtermiosの設定
func SettingTermios() {
	termios.Tcgetattr(uintptr(syscall.Stdin), &Editor.defaultTtystate)
	ttystate := Editor.defaultTtystate
	setRawMode(&ttystate)
}

// 非カノニカルモードに設定する
func setRawMode(attr *unix.Termios) {
	attr.Iflag &^= syscall.BRKINT | syscall.ICRNL | syscall.INPCK | syscall.ISTRIP | syscall.IXON
	attr.Cflag &^= syscall.CSIZE | syscall.PARENB
	attr.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.IEXTEN
	attr.Cc[syscall.VMIN] = 1
	attr.Cc[syscall.VTIME] = 0
	termios.Tcsetattr(uintptr(syscall.Stdin), termios.TCSANOW, attr)
}

// ターミナル属性をリセットする
func resetRawMode(attr *unix.Termios) {
	termios.Tcsetattr(uintptr(syscall.Stdin), termios.TCSANOW, attr)
}