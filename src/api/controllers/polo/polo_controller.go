package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

func Polo(c *gin.Context) {
	// let cloud know your webserver is ready for handling traffic
	c.String(http.StatusOK, polo)
}
