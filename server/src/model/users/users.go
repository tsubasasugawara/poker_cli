package users

import (
	"database/sql"
	"crypto/sha256"
	"os"
	"errors"

	_ "github.com/lib/pq"
)

const (
	dbtype = "postgres"

	OK = 0
	IllegalName = 1
	IllegalPassword = 2
	NotOpening = 3
	NotPreparing = 4
	NotExecution = 5
)

// ハッシュ化関数
func hash(s string) string {
    r := sha256.Sum256([]byte(s))
    return string(r[:])
}

type Users struct {
	dbUrl string
}

func NewUsers() *Users {
	users := Users{}
	users.dbUrl = "host=" + os.Getenv("POSTGRES_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=disable"
	return &users
}

/*
 * ユーザ登録
 * @{param} name string
 * @{param} pasword string
 * @{result} int : 成功したときは0、失敗したときは0以外を返す
 * @{result} error
*/
func(u *Users) Regist(name, password string) (int, error) {
	// バリデーション
	if ValidateName(name) == false {
		return IllegalName, errors.New("Illegal name.")
	}
	if ValidatePassword(password) == false {
		return IllegalPassword, errors.New("Illegal password.")
	}

	db, err := sql.Open(dbtype, u.dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (name, password) VALUES ($1, $2)")
	if err != nil {
		return NotPreparing, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(name, hash(password)).Scan(&id)
	if err != nil {
		return NotExecution, err
	}

	return OK, nil
}

/*
 * ログイン
 * @{param} name string
 * @{param} password string
 * @{result} int : 成功したときは0、失敗したときは0以外を返す
 * @{result} error
*/
func(u *Users) Select(name, password string) (int, error) {
	db, err := sql.Open(dbtype, u.dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id FROM users WHERE name = $1 AND password = $2")
	if err != nil {
		return NotPreparing, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(name, hash(password)).Scan(&id)
	if err != nil {
		return NotExecution, err
	}

	return OK, nil
}

/*
 * ユーザー削除
 * @{param} name string
 * @{param} password string
 * @{result} int : 成功したときは0、失敗したときは0以外を返す
 * @{result} error
*/
func(u *Users) Delete(name, password string) (int, error) {
	db, err := sql.Open(dbtype, u.dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE name = $1 AND password = $2")
	if err != nil {
		return NotPreparing, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, hash(password))
	if err != nil {
		return NotExecution, err
	}

	return OK, nil
}

/*
 * ユーザ情報変更
 * @{param} name string
 * @{param} password string
 * @{result} int : 成功したときは0、失敗したときは0以外を返す
 * @{result} error
*/
func(u *Users) Update(name, password string) (int, error){
	// バリデーション
	if ValidateName(name) == false {
		return IllegalName, errors.New("Illegal name.")
	}
	if ValidatePassword(password) == false {
		return IllegalPassword, errors.New("Illegal password.")
	}

	db, err := sql.Open(dbtype, u.dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET name = $1, password = $2 WHERE name = $1 AND password = $2")
	if err != nil {
		return NotPreparing, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, hash(password))
	if err != nil {
		return NotExecution, err
	}

	return OK, nil
}
