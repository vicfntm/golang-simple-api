package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vicfntm/splitshit2/src/handlers/exceptionhandlers"
)

const USED_MUST_BE_FOUND = "user must be found in db"
const USER_DOES_NOT_EXIST = "user not found"
const NOT_ALLOWED = "operation not allowed"

func (h *Handler) userSuperAdminMiddleware(c *gin.Context) {
	userId, exist := c.Get("user")
	if !exist {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, USED_MUST_BE_FOUND)
	}

	dbuser, error := h.services.Authorization.FindCustomer(userId)

	if error != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusNotFound, USER_DOES_NOT_EXIST)
	}

	adminRole := viper.GetStringMapString("ROLES")

	if dbuser.Role != adminRole["sadmin"] {
		exceptionhandlers.NewErrorResponse(c, http.StatusForbidden, NOT_ALLOWED)
	}
	c.Set("role", dbuser.Role)

}
