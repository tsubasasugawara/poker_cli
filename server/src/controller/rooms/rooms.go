package rooms

import (
	"log"
	"net/http"
	"poker/model/rooms"

    "github.com/gin-gonic/gin"
)

type Room struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	RoomId string `json:"roomId"`
}

func CreateRoom(c *gin.Context) {
	var room Room
	c.BindJSON(&room)

	var err error
	room.RoomId, err = rooms.Insert(room.UserId, room.Password)
	if err != nil {
		log.Println(err)
	}

	var msg = "success"
	if room.RoomId == "" {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
		"roomId": room.RoomId,
	})
}

func ParticipateRoom(c *gin.Context) {
	var room Room
	c.BindJSON(&room)

	statusCode, err := rooms.Select(room.RoomId, room.Password)
	if err != nil {
		log.Println(err)
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}

// func ExitRoom(c *gin.Context) {
// }
