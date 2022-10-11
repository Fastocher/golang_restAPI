package restapp

type User struct {
	//для того чтобы get из базы работал
	//добавляю тег для поля id
	Id       int    `json:"=" db "id"`
	Name     string `json: "name" binding:"required"`
	Username string `json: "username" binding:"required"`
	Password string `json: "password" binding:"required"`
}
