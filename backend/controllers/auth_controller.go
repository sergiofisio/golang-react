package controllers

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergiofisio/golang-react/config"
	"github.com/sergiofisio/golang-react/models"
	"github.com/sergiofisio/golang-react/services"
	"github.com/sergiofisio/golang-react/utils"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		msg := "Dados inválidos: " + err.Error()
		if errors.Is(err, io.EOF) {
			msg = "Corpo da requisição JSON está vazio"
		}
		utils.SendErrorResponse(c, http.StatusBadRequest, msg, "Register")
		return
	}

	hashedPassword, err := services.HashPassword(input.Password)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Erro ao gerar hash da senha", "Register")
		return
	}

	user := models.Users{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
		Active:   true,
		Role:     models.User,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Usuario com username e/ou email já cadastrado", "Register")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
