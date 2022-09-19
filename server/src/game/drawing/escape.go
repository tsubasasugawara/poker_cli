package drawing

import (
	"fmt"
)

/*
 * エスケープ処理を行う
 * @{param} str string : エスケープしたい文字列
 * @{result} string : エスケープした文字列
 */
func Escape(str string) (string) {
	return fmt.Sprintf("%#v", str)
}
