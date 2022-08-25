package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicfntm/splitshit2/src/handlers/exceptionhandlers"
	"github.com/vicfntm/splitshit2/src/models"
)

func (h *Handler) createGroup(c *gin.Context) {

	var input models.Groups

	if err := c.BindJSON(&input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, exist := c.Get("user")

	if !exist {
		log.Fatalf("user not found")
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"userId":  user,
	})
}

func (h *Handler) availableGroups(c *gin.Context) {

}

func (h *Handler) joinGroup(c *gin.Context) {

}

func (h *Handler) grabShit(c *gin.Context) {

}
