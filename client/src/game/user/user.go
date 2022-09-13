package user

type User struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type Message struct {
	Id string `json:"id"`
	Status string `json:"statusMessage"`
}
