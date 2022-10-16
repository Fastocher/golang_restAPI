package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "userId"
)

// прослойка для парсинга токенов из запроса и предоставления доступа к группе ендпоинтов /api
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	//разделяем хедер по пробелам
	//должны получить массив из двух элементов
	//
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	//парсинг токена
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Записываем id пользователя в gin.Context для того
	// чтобы иметь доступ к этому значению в последующих обработчиках
	// которые вызываются после данной прослойки
	c.Set(userContext, userId)

}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userContext)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idint, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id invalid type")
		return 0, errors.New("user id invalid type")
	}

	return idint, nil

}
