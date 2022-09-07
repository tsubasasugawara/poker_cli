package rooms

import (
	"log"
	"net/http"
	"poker/model/rooms"
	"poker/model/participants"

    "github.com/gin-gonic/gin"
)

type Room struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	RoomId string `json:"roomId"`
}

/*
 * ルーム作成
 * @{param} c *gin.Context
 * @{request json} "userId" : "作成者のユーザID"
 * @{request json} "password" : "ルームに設定したいパスワード"
 * @{response} json {"statusMessage" : "success" or "failure"}
 * @{response} json {"roomId" : "ルームID(36)" or ""}
 * @{response} json {"password" : password}
*/
func CreateRoom(c *gin.Context) {
	var room Room
	c.BindJSON(&room)

	var err error
	room.RoomId, err = rooms.Insert(room.UserId, room.Password)
	if err != nil {
		log.Println(err)
		return
	}

	var msg = "success"
	if room.RoomId == "" {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
		"roomId": room.RoomId,
		"password": room.Password,
	})
}

/*
 * ルームに入る
 * @{param} c *gin.Context
 * @{request json} "userId" : "ユーザID"
 * @{request json} "roomId" : "ルームID"
 * @{request json} "password" : "ルームのパスワード"
 * @{response} json {"statusMessage" : "success" or "failure"}
 * @{response} json {"playerCount" : 参加者数}
*/
func ParticipateRoom(c *gin.Context) {
	var room Room
	c.BindJSON(&room)

	cnt, err := participants.CountParticipants(room.RoomId)
	if err != nil {
		log.Println(err)
		return
	}
	if cnt >= 2 {
		return
	}

	statusCode, err := rooms.Select(room.RoomId, room.Password)
	if err != nil {
		log.Println(err)
		return
	}

	statusCode, err = participants.Insert(room.RoomId, room.UserId)
	if err != nil {
		log.Println(err)
		return
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
		"playerCount": cnt + 1,
	})
}

/*
 * ルームから退出
 * @{param} c *gin.Context
 * @{request json} "userId" : "ユーザID"
 * @{request json} "roomId" : "ルームID"
 * @{response} json {"statusMessage" : "success" or "failure"}
*/
func ExitRoom(c *gin.Context) {
	var room Room
	c.BindJSON(&room)

	statusCode, err := participants.Delete(room.RoomId, room.UserId)
	if err != nil {
		log.Println(err)
		return
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}
