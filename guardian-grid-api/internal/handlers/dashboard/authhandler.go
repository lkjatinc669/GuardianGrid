// internal/handlers/auth.go
package dashboard

import (
	"guardian-grid-api/internal/auth"
	"guardian-grid-api/internal/db"
	"guardian-grid-api/internal/models"

	"github.com/gin-gonic/gin"
)

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Result struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Register Endpoint

func Register(c *gin.Context) {
	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, Result{
			Error:   true,
			Message: "Invalid Input",
			Data:    "none",
		})
		return
	}

	// 🔹 Username length check
	if len(input.Username) < 5 {
		c.JSON(400, Result{
			Error:   true,
			Message: "Username must be at least 5 characters",
			Data:    "none",
		})
		return
	}

	// 🔹 Password length check
	if len(input.Password) < 8 {
		c.JSON(400, Result{
			Error:   true,
			Message: "Password must be at least 8 characters",
			Data:    "none",
		})
		return
	}

	// 🔹 Check uniqueness
	var existing models.User
	err := db.SQLite.Where("username = ?", input.Username).First(&existing).Error
	if err == nil {
		c.JSON(400, Result{
			Error:   true,
			Message: "Username already exists",
			Data:    "none",
		})
		return
	}

	// 🔹 Create user
	user := models.User{
		Username: input.Username,
		Password: auth.HashPassword(input.Password),
	}

	db.SQLite.Create(&user)

	c.JSON(200, Result{
		Error:   false,
		Message: "User created",
		Data:    "none",
	})
}

// Login Endpoint

func Login(c *gin.Context) {
	var input AuthInput

	// 🔹 Validate JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, Result{
			Error:   true,
			Message: "Invalid Input",
			Data:    "none",
		})
		return
	}

	// 🔹 Basic username validation
	if input.Username == "" {
		c.JSON(400, Result{
			Error:   true,
			Message: "Username required",
			Data:    "none",
		})
		return
	}

	// 🔹 Basic password validation
	if input.Password == "" {
		c.JSON(400, Result{
			Error:   true,
			Message: "Password required",
			Data:    "none",
		})
		return
	}

	// 🔹 Find user
	var user models.User
	err := db.SQLite.Where("username = ?", input.Username).First(&user).Error
	if err != nil {
		c.JSON(401, Result{
			Error:   true,
			Message: "Invalid credentials",
			Data:    "none",
		})
		return
	}

	// 🔹 Check password
	if !auth.CheckPassword(input.Password, user.Password) {
		c.JSON(401, Result{
			Error:   true,
			Message: "Invalid credentials",
			Data:    "none",
		})
		return
	}

	// 🔹 Generate token
	token := auth.GenerateToken(user.ID)

	c.JSON(200, Result{
		Error:   false,
		Message: "Login Successful",
		Data:    token,
	})
}
