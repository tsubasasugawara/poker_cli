package participants

import (
	"database/sql"
	"errors"

	"poker/model"

	_ "github.com/lib/pq"
)

/*
 * ルームに参加者が何人いるのかを確認
 * @{param} roomId string
 * @{result} int 成功したときはプレイヤー数, 失敗したときは0未満のステータスコード
 * @{result} error
*/
func CountParticipants(roomId string) (int, error) {
	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	var count int
	err = db.QueryRow(
		"SELECT COUNT(*) FROM participants WHERE room_id = $1",
		roomId,
	).Scan(&count)
	if err != nil {
		return model.NotExecution, err
	}

	return count, nil
}

/*
 * ルームに参加
 * @{param} roomId string
 * @{param} userId string
 * @{result} int 成功したときは0, 失敗したときは0未満のステータスコード
 * @{result} error
*/
func Insert(roomId, userId string) (int, error) {
	// バリデーション
	if model.ValidateRoomId(roomId) == false {
		return model.IllegalRoomId, errors.New("Illegal room id.")
	}
	if model.ValidateUserId(userId) == false {
		return model.IllegalUserId, errors.New("Illegal user id.")
	}

	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO participants(room_id, player_id) VALUES ($1, $2)",
		roomId,
		userId,
	)
	if err != nil {
		return model.NotExecution, err
	}

	return model.OK, nil
}

/*
 * ルームから退出
 * @{param} roomId string
 * @{param} userId string
 * @{result} int 成功したときは0, 失敗したときは0未満のステータスコード
 * @{result} error
*/
func Delete(roomId, userId string) (int, error) {
	db, err := sql.Open(model.DBtype, model.DBUrl)
	if err != nil {
		return model.NotOpening, err
	}
	defer db.Close()

	_, err = db.Exec(
		"DELETE FROM participants WHERE room_id = $1 AND player_id = $2",
		roomId,
		userId,
	)
	if err != nil {
		return model.NotExecution, err
	}

	return model.OK, nil
}

