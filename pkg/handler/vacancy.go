package handler

import (
	"alnshine/CRUD_FOR_BAL"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create vacancy
// @Security ApiKeyAuth
// @Tags vacancies
// @Description create vacancy
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body CRUD_FOR_BAL.Vacancy true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/vacancy [post]
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

// @Summary Get All Vacancies
// @Security ApiKeyAuth
// @Tags vacancies
// @Description get all vacancies
// @ID get-all-vacancies
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/vacancy [get]
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

// @Summary Get Vacancy By Id
// @Security ApiKeyAuth
// @Tags vacancies
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} CRUD_FOR_BAL.Vacancy
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/vacancy/:id [get]
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
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input CRUD_FOR_BAL.UpdateVac
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{"ok"})
}
func (h *Handler) deleteVacancy(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Vacancy.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{
		"ok",
	})
}
