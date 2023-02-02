package root

import (
	"net/http"

	"github.com/Kkmikaze/go-rest-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	response.Json(context, http.StatusOK, "", nil)
}
