package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createVacancy(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getAllVacancy(c *gin.Context) {

}
func (h *Handler) getVacancyById(c *gin.Context) {

}
func (h *Handler) updateVacancy(c *gin.Context) {

}
func (h *Handler) deleteVacancy(c *gin.Context) {

}
