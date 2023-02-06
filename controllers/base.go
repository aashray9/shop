package controllers

import (
	"lms/common"
	"lms/database"
	"lms/entity"
	"lms/models"
	"lms/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var log = common.Loggers()

func BaseRoutes(router *gin.RouterGroup) {
	router.POST("/get-token", GetToken)
	router.GET("/check", CheckConfig)
}
func GetToken(c *gin.Context) {
	var request entity.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Debug().Msg(err.Error())
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}
	log.Info().Msg(request.Email)

	db := database.GetConnection()
	var user models.CrmUsers
	result := db.Select("email", "name", "id").Where("email", request.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, common.NewError("loggin Error", result.Error))
		return
	}

	// bytePassword := []byte(request.Password)
	// byteHashedPassword := []byte(user.Password)
	// err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	// if err != nil {
	// 	c.JSON(http.StatusForbidden, dtobjects.DetailedErrors("login", errors.New("invalid credential")))
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   user.GenerateJwtToken(),
		"user_id": user.Id,
	})
}

func CheckConfig(c *gin.Context) {
	leadSource := services.LeadSource()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    leadSource,
	})
}
