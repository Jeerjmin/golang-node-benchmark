package validator

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		Validate: validator.New(),
	}
}

func ValidateRequest[T any](v *Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := v.Validate.Struct(req); err != nil {
			errors := err.(validator.ValidationErrors)
			errorMessages := make(map[string]string)

			for _, e := range errors {
				errorMessages[e.Field()] = e.Error()
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			c.Abort()
			return
		}

		c.Set("validatedRequest", req)
		c.Next()
	}
}

func (v *Validator) ValidateIDParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if _, err := strconv.Atoi(id); err != nil || id == "0" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
			c.Abort()
			return
		}

		c.Next()
	}
}
