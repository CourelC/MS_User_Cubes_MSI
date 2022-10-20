package user

import (
	"MS_User_Cubes_MSI/controllers"
	errorModels "MS_User_Cubes_MSI/errors"
	"MS_User_Cubes_MSI/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// NewUser godoc
// @Tags user
// @Summary Create a new user
// @Description Create a new user
// @Accept  json
// @Produce  json
// @Param data body NewUserRequest true "body data"
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/ [post]
func NewUser(c *gin.Context) {
	var request NewUserRequest

	if err := controllers.BindJSON(c, &request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}
	user := models.User{
		Name:         request.Name,
		Mail:         request.Mail,
		Password:     request.Password,
		CreationDate: time.Now(),
		IsActive:     false,
		IsValided:    false,
	}

	err := models.CreateUser(&user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUser godoc
// @Tags user
// @Summary Get Users
// @Description Get all Users on the system
// @Success 200 {object} []models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	err := models.GetUsers()
	if err != nil {
		appError := errorModels.NewAppErrorWithType(errorModels.UnknownError)
		_ = c.Error(appError)
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserById godoc
// @Tags user
// @Summary Get detail User
// @Description Get all Users on the system
// @Param user_id path int true "id of user"
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/{user_id} [get]
func GetUserById(c *gin.Context) {
	var user models.User
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	err = models.GetUserByID(userId)
	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, user)
}

// ValideUser godoc
// @Tags user
// @Summary Valide User
// @Description Validation user by admin
// @Param user_id path int true "id of user"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/{user_id}/valided [put]
func ValidedUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	err = models.ValidedUser(userId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "resource valided successfully"})

}

// InactiveUser godoc
// @Tags user
// @Summary Inactive User
// @Description Inactivate user by admin
// @Param user_id path int true "id of user"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/{user_id}/inactive [put]
func InactiveUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	err = models.ValidedUser(userId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "resource inactived successfully"})

}

// AuthUser godoc
// @Tags user
// @Summary Authentification User
// @Description Inactivate user by admin
// @Param user_id path int true "id of user"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/auth [get]
func AuthUser(c *gin.Context) {
	var request AuthUserRequest

	if err := controllers.BindJSON(c, &request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	user := models.User{
		Mail:     request.Mail,
		Password: request.Password,
	}

	err := models.AuthUser(user.Mail, user.Password)
	if err != nil {
		appError := errorModels.NewAppErrorWithType(errorModels.UnknownError)
		_ = c.Error(appError)
		return
	}
	c.JSON(http.StatusOK, user)
}
