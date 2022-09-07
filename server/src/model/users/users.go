package users

import (
	"database/sql"
	"errors"

	"poker/lib/uuid"
	"poker/model"

	_ "github.com/lib/pq"
)

/*
 * ユーザ登録
 * @{param} name string
 * @{param} pasword string
 * @{result} int : 成功したときは0、失敗したときは0以外を返す
 * @{result} error
*/
func Insert(name, password string) (int, error) {
	// バリデーション
	if model.ValidateName(name) == false {
		return model.IllegalName, errors.New("Illegal name.")
	}
	if model.ValidatePassword(password) == false {
		return model.IllegalPassword, errors.New("Illegal password.")
	}

	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO users (id, name, password) VALUES ($1, $2, $3)",
		uuid.Generate(),
		name,
		model.Hash(password),
	)
	if err != nil {
		return model.NotExecution, err
	}

	return model.OK, nil
}

/*
 * select
 * @{param} name string
 * @{param} password string
 * @{result} string : 成功したときはid、失敗したときは空
 * @{result} error
*/
func Select(name, password string) (string, error) {
	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string
	err = db.QueryRow(
		"SELECT id FROM users WHERE name = $1 AND password = $2",
		name,
		model.Hash(password),
		).Scan(&id)

	if err != nil {
		return "", err
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
	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	id, err := Select(name, password)
	if err != nil {
		return model.NotExecution, err
	}

	// roomsのuser_id_created_roomが外部キーに設定されているため、先に削除する
	_, err = db.Exec(
		"DELETE FROM rooms WHERE user_id_created_room = $1",
		id,
	)
	if err != nil {
		return model.NotExecution, err
	}

	_, err = db.Exec(
		"DELETE FROM users WHERE id = $1",
		id,
	)
	if err != nil {
		return model.NotExecution, err
	}

	return model.OK, nil
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
	if model.ValidateName(oldName) == false {
		return model.IllegalName, errors.New("Illegal name.")
	}
	if model.ValidatePassword(oldPassword) == false {
		return model.IllegalPassword, errors.New("Illegal password.")
	}
	if model.ValidateName(newName) == false {
		return model.IllegalName, errors.New("Illegal name.")
	}
	if model.ValidatePassword(newPassword) == false {
		return model.IllegalPassword, errors.New("Illegal password.")
	}

	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec(
		"UPDATE users SET name = $1, password = $2 WHERE name = $3 AND password = $4",
		newName,
		model.Hash(newPassword),
		oldName,
		model.Hash(oldPassword),
	)
	if err != nil {
		return model.NotExecution, err
	}

	return model.OK, nil
}


func UpdateAccessDate(id string) (int, error) {
	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec(
		"UPDATE users SET access_date = DEFAULT WHERE id = $1",
		id,
	)
	if err != nil {
		return model.NotExecution, err
	}

	return model.OK, nil
}
