package Services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(c *gin.Context, status int, message , data interface {}) {
	c.AbortWithStatusJSON(status, gin.H{"message": message, "data": data})
}

func ErrorResponse(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"message": message})
}

//200
func Success(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusOK, message, data)
}
//201
func Created(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusCreated, message, data)
}
//204
func NoContent(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusNoContent, message, data)
}
//302
func RedirectLink(c *gin.Context, originalId string) {
	c.Redirect(http.StatusTemporaryRedirect, originalId)
}
//400
func BadRequest(c *gin.Context, message interface{}) {
	ErrorResponse(c, http.StatusBadRequest, message)
}
//401
func Unauthorized(c *gin.Context, message interface{}) {
	ErrorResponse(c, http.StatusUnauthorized, message)
}
//406
func NotAcceptable(c *gin.Context, message interface{}) {
	ErrorResponse(c, http.StatusNotAcceptable, message)
}
//409
func AlreadyExists(c *gin.Context, message interface{}) {
	ErrorResponse(c, http.StatusConflict, message)
}
//422
func ValidationError(c *gin.Context, message interface{}) {
	ErrorResponse(c, http.StatusUnprocessableEntity, message)
}
//500
func InternalError(c *gin.Context, message interface{}) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}