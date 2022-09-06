package users

import (
	"database/sql"
	"crypto/sha256"
	"os"
	"errors"
	"encoding/hex"

	_ "github.com/lib/pq"
)

const (
	dbtype = "postgres"

	OK = 0
	IllegalName = -1
	IllegalPassword = -2
	NotOpening = -3
	NotExecution = -4
)

var dbUrl = os.Getenv("DB_URL")

// ハッシュ化関数
func hash(s string) string {
    r := sha256.Sum256([]byte(s))
    return hex.EncodeToString(r[:])
}

/*
 * ユーザ登録
 * @{param} name string
 * @{param} pasword string
 * @{result} int : 成功したときは0、失敗したときは0以外を返す
 * @{result} error
*/
func Insert(name, password string) (int, error) {
	// バリデーション
	if ValidateName(name) == false {
		return IllegalName, errors.New("Illegal name.")
	}
	if ValidatePassword(password) == false {
		return IllegalPassword, errors.New("Illegal password.")
	}

	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", name, hash(password))
	if err != nil {
		return NotExecution, err
	}

	return OK, nil
}

/*
 * select
 * @{param} name string
 * @{param} password string
 * @{result} int : 成功したときはid、失敗したときは0以外を返す
 * @{result} error
*/
func Select(name, password string) (int, error) {
	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	var id int
	err = db.QueryRow("SELECT id FROM users WHERE name = $1 AND password = $2", name, hash(password)).Scan(&id)
	if err != nil {
		return NotExecution, err
	}

	return id, nil
}

/*
 * ユーザー削除
 * @{param} name string
 * @{param} password string
 * @{result} int : 成功したときは0、失敗したときは0以外を返す
 * @{result} error
*/
func Delete(name, password string) (int, error) {
	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE name = $1 AND password = $2", name, hash(password))
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
func Update(oldName, oldPassword, newName, newPassword string) (int, error){
	// バリデーション
	if ValidateName(oldName) == false {
		return IllegalName, errors.New("Illegal name.")
	}
	if ValidatePassword(oldPassword) == false {
		return IllegalPassword, errors.New("Illegal password.")
	}

	if ValidateName(newName) == false {
		return IllegalName, errors.New("Illegal name.")
	}
	if ValidatePassword(newPassword) == false {
		return IllegalPassword, errors.New("Illegal password.")
	}

	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET name = $1, password = $2 WHERE name = $3 AND password = $4", newName, hash(newPassword), oldName, hash(oldPassword))
	if err != nil {
		return NotExecution, err
	}

	return OK, nil
}
