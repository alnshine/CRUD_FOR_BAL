package handler

import (
	"alnshine/CRUD_FOR_BAL"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createVacancy(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input CRUD_FOR_BAL.Vacancy
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Vacancy.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []CRUD_FOR_BAL.Vacancy `json:"data"`
}

func (h *Handler) getAllVacancy(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	vac, err := h.services.Vacancy.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: vac,
	})
}
func (h *Handler) getVacancyById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	vac, err := h.services.Vacancy.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vac)
}
func (h *Handler) updateVacancy(c *gin.Context) {

}
func (h *Handler) deleteVacancy(c *gin.Context) {

}
