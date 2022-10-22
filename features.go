package restapp

type UsersMessages struct {
	Id        int
	UserId    int
	MessageId int
}

// теги db для возможности выборки из базы
type Message struct {
	Id      int    `json:"id" db:"id"`
	Message string `json:"message" db:"message" binding:"required"`
}
