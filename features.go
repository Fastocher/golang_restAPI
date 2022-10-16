package restapp

type UsersMessages struct {
	Id        int
	UserId    int
	MessageId int
}

type Message struct {
	Id      int    `json:" id"`
	Message string `json: "message" binding:"required"`
}
