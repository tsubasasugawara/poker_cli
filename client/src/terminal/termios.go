package terminal

import (
	"syscall"
	"golang.org/x/sys/unix"

	"github.com/pkg/term/termios"
)

// 起動時のtermiosの設定
func SettingTermios(defaultTtyState *unix.Termios) {
	termios.Tcgetattr(uintptr(syscall.Stdin), defaultTtyState)
	setRawMode(defaultTtyState)
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
