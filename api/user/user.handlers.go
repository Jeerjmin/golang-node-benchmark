package user

import (
	"go-auth/internal/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
			Role     string `json:"role"`
		}
		// Попробуйте распарсить JSON из тела запроса в req
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Чистый SQL запрос для вставки данных
		query := `
			INSERT INTO "users" ("name", "email", "role", "password")
			VALUES (?, ?, ?, ?)
			RETURNING "id", "name", "email", "role";
		`

		// Переменная для хранения данных нового пользователя
		var user struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
			Role  string `json:"role"`
		}

		// Выполнение SQL-запроса через GORM
		if err := db.Raw(query, req.Name, req.Email, req.Role, req.Password).Scan(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Возвращаем данные нового пользователя в JSON формате
		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}

func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")

		var user models.User
		if err := db.First(&user, userId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Чистый SQL-запрос для выборки пользователей
		query := `
			SELECT "id", "name", "email", "role", "password"
			FROM "users"
			LIMIT 20;
		`

		// Переменная для хранения данных пользователей
		var users []models.User

		// Выполнение SQL-запроса через GORM
		if err := db.Raw(query).Scan(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Возвращаем данные пользователей в JSON формате
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet("validatedRequest").(UpdateUserRequest)

		userId := c.Param("id")

		var user models.User
		if err := db.First(&user, userId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		user.Name = req.Name

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})

	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")

		if err := db.Delete(&models.User{}, userId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
