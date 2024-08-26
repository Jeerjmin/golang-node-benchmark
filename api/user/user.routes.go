package user

import (
	"go-auth/api/validator"
	database "go-auth/internal/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *database.Database) *gin.Engine {
	v := validator.NewValidator()

	router.POST("/users", CreateUser(db.DB))
	router.GET("/users/:id", v.ValidateIDParam(), GetUser(db.DB))
	router.GET("/users", GetUsers(db.DB))
	router.PUT("/users/:id", v.ValidateIDParam(), validator.ValidateRequest[UpdateUserRequest](v), UpdateUser(db.DB))
	router.DELETE("/users/:id", v.ValidateIDParam(), DeleteUser(db.DB))

	return router
}
