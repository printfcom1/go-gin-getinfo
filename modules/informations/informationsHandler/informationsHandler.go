package informationsHandler

import (
	"net/http"

	"github.com/gin-api/modules/informations/informationsService"
	"github.com/gin-gonic/gin"
)

type informationsHandler struct {
	informationsHand informationsService.InformationsService
}

func NewInformationsHandler(infoHand informationsService.InformationsService) informationsHandler {
	return informationsHandler{informationsHand: infoHand}
}

func (h *informationsHandler) GetInformations(c *gin.Context) {
	info, err := h.informationsHand.GetInformationsService()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"code":   200,
		"data":   info,
	})
}
