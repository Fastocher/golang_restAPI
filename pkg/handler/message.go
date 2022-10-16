package handler

import (
	"net/http"

	"github.com/Fastocher/restapp"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMessage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input restapp.Message
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//реализация метода создания сообщения

	id, err := h.services.Message.CreateMessage(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//добавление тела ответа

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllMessages(c *gin.Context) {

}

func (h *Handler) getMessageById(c *gin.Context) {

}

func (h *Handler) updateMessage(c *gin.Context) {

}

func (h *Handler) deleteMessage(c *gin.Context) {

}
