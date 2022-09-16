package room

type Room struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	RoomId string `json:"roomId"`
}

type Message struct {
	Status string `json:"statusMessage"`
}
