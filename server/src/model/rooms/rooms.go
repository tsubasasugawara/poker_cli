package rooms

import (
	"poker/model"
	"poker/lib/uuid"

	_ "github.com/lib/pq"
)

/*
 * ルーム登録
 * @{param} userid string
 * @{param} pasword string ルームのパスワード
 * @{result} string 成功したときはルームID、失敗したときは空
 * @{result} error
*/
func Insert(userid, password string) (string, error) {
	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return "", err
	}
	defer db.Close()

	uuid := uuid.Generate()
	_, err = db.Exec(
		"INSERT INTO rooms (id, user_id_created_room, password) VALUES ($1, $2, $3)",
		uuid,
		model.Hash(password),
		userid,
	)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

/*
 * ルームへアクセス
 * @{param} roomId string
 * @{param} password string
 * @{result} int 成功したときは0、失敗したときは0以外
 * @{result} error
*/
func Select(roomId, password string) (int, error) {
	db, err := sql.Open(dbtype, dbUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	uuid := uuid.Generate()
	_, err = db.Exec(
		"SELECT id FROM rooms WHERE id = $1 and password = $2",
		roomId,
		model.Hash(password),
	)
	if err != nil {
		return model.NotExecution, err
	}

	return uuid, nil
}
