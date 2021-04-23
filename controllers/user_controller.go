package controllers

import (
	"net/http"
	"project/configs"
	"project/middleware"
	"project/models"

	"github.com/labstack/echo"
)

func LoginUsersController(c echo.Context) error {
	var userInput models.UserRequest
	c.Bind(&userInput)

	var userDB models.Users

	err := configs.DB.Where("email = ? AND password = ?", userInput.Email, userInput.Password).Find(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	token, err := middleware.GenerateToken(int(userDB.ID), userDB.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error jwt",
		})
	}

	var userResponse = models.UsersDBResponse{
		ID:    userDB.ID,
		Name:  userDB.Name,
		Email: userDB.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, models.UserResponseSingle{
		Code:    http.StatusOK,
		Message: "Success get data user",
		Status:  "success",
		Data:    userResponse,
	})
}

func GetUsersController(c echo.Context) error {
	// categoryId := c.QueryParam("categoryId")
	// page := c.QueryParam("page")
	// userId, _ := strconv.Atoi(c.Param("userId"))

	var users []models.Users
	err := configs.DB.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	// userId dari JWT
	userId := middleware.ExtractUserIdFromJWT(c)

	return c.JSON(http.StatusOK, models.UserResponse{
		Code:    userId,
		Message: "Success get data user",
		Status:  "success",
		Data:    users,
	})
}

func CreateUsersController(c echo.Context) error {

	// email := c.FormValue("email")
	// name := c.FormValue("name")

	var userInput models.UserRequest
	c.Bind(&userInput)

	var userDB models.Users
	userDB.Name = userInput.Name
	userDB.Email = userInput.Email
	userDB.Password = userInput.Password
	// userDB.School = models.School{Name: "Alta"}

	err := configs.DB.Save(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	token, err := middleware.GenerateToken(int(userDB.ID), userDB.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error jwt",
		})
	}

	var userResponse = models.UsersDBResponse{
		ID:    userDB.ID,
		Name:  userDB.Name,
		Email: userDB.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, models.UserResponseSingle{
		Code:    http.StatusOK,
		Message: "Success get data user",
		Status:  "success",
		Data:    userResponse,
	})
}
