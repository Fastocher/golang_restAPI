package handler

import (
	"net/http"
	"strconv"

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

type getAllMessagesResponse struct {
	Data []restapp.Message `json: "data"`
}

func (h *Handler) getAllMessages(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	messages, err := h.services.Message.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllMessagesResponse{
		Data: messages,
	})
}

func (h *Handler) getMessageById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	message, err := h.services.Message.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, message)
}

func (h *Handler) updateMessage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid messageId")
	}

	var input restapp.UpdateMessageInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteMessage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	err = h.services.Message.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
