package model

import (
	"database/sql"
	"crypto/sha256"
	"os"
	"errors"

	_ "github.com/lib/pq"
)

const (
	dbtype := "postgres"
	dbUrl := "host=" + os.Getenv("POSTGRES_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=disable"
)

// ハッシュ化関数
func hash(s string) string {
    r := sha256.Sum256([]byte(s))
    return string(r[:])
}

/*
 * ユーザ登録
 * @{param} name string
 * @{param} pasword string
 * @{result} id int : 成功したときはid、失敗したときは-1を返す
 * @{result} error
*/
func Regist(name, password string) (int, error) {
	// バリデーション
	if ValidateName(name) == false {
		return -1, errors.New("Illegal name.")
	}
	if ValidatePassword(password) == false {
		return -1, errors.New("Illegal password.")
	}

	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (name, password) VALUES($1, $2)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var id int
	err := stmt.QueryRow(name, hash(password)).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

/*
 * ログイン
 * @{param} name string
 * @{param} password string
 * @{result} id int : カラムが見つかったらidを返す
 * @{result} error
*/
func Select(name, password string) (int, error) {
	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id FROM users WHERE name = $1 AND password = $2")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var id int
	err := stmt.QueryRow(name, hash(password)).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

/*
 * ユーザー削除
 * @{param} name string
 * @{param} password string
 * @{result} error
*/
func Delete(name, password string) error {
	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE name = $1 AND password = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err := stmt.Exec(name, hash(password))
	if err != nil {
		return err
	}

	return nil
}

/*
 * ユーザ情報変更
 * @{param} name string
 * @{param} password string
 * @{result} id int : カラムが見つかったらidを返す
 * @{result} error
*/
func Update(name, password string) error{
	// バリデーション
	if ValidateName(name) == false {
		return -1, errors.New("Illegal name.")
	}
	if ValidatePassword(password) == false {
		return -1, errors.New("Illegal password.")
	}

	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET name = $1, password = $2 WHERE name = $1 AND password = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err := stmt.Exec(name, hash(password))
	if err != nil {
		return err
	}

	return nil
}
