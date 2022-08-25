package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vicfntm/splitshit2/src/handlers/exceptionhandlers"
	"github.com/vicfntm/splitshit2/src/models"
)

func (h *Handler) signIn(c *gin.Context) {
	var input models.Customer

	if err := c.BindJSON(&input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.services.Authorization.LoginUser(input)
	if err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) signUp(c *gin.Context) {

	var input models.Customer

	if err := c.BindJSON(&input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateCustomer(input)

	if err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) assignRole(c *gin.Context) {

	var input models.RoleCustomer
	data := viper.GetStringMapString("ROLES")
	if err := c.BindJSON(&input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	customer, err := h.services.Authorization.AssignRole(input.Id, strings.ToLower(data[input.Role]))

	if err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	reverted := revertMap(data)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":   customer.Id,
		"role": reverted[customer.Role],
	})

}

func revertMap(data map[string]string) map[string]string {
	res := map[string]string{}
	for k, elem := range data {
		res[elem] = k
	}

	return res
}
