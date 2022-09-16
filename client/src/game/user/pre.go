package user

import (
	"bufio"
	"fmt"
)

func Pre(scanner *bufio.Scanner) (string) {
	running := true
	var userId string

	for running {
		fmt.Println("Please log in (l) or sign up (s).")
		scanner.Scan()
		switch scanner.Text() {
		// ログイン
		case "l":
			fmt.Print("Name : ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Password : ")
			scanner.Scan()
			password := scanner.Text()

			// ログイン処理
			res := login(name, password)

			if res != "" {
				userId = res
				running = false
			}

		// サインアップ
		case "s":
			fmt.Print("Name : ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Password : ")
			scanner.Scan()
			password := scanner.Text()

			// 登録処理
			err := signup(name, password)
			if err != nil{
				fmt.Println("Sorry, it didn't work.")
				continue
			}

			fmt.Println("Please login.")
		}
	}

	return userId
}
