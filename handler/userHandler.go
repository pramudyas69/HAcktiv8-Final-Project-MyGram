package handler

import (
	"MyGramHacktiv8/dto"
	"MyGramHacktiv8/entity"
	"MyGramHacktiv8/pkg/helpers"
	"MyGramHacktiv8/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserRestHandler(userService service.UserService) *userRestHandler {
	return &userRestHandler{userService: userService}
}

func (u *userRestHandler) Login(c *gin.Context) {
	var userRequest dto.LoginRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	token, err2 := u.userService.Login(&userRequest)

	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}
	c.JSON(http.StatusCreated, token)
}

func (u *userRestHandler) Register(c *gin.Context) {
	var userRequest dto.RegisterRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		// ! TODO: JSON bind not working
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("User =>", &userRequest)
	successMessage, err2 := u.userService.Register(&userRequest)

	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Message(),
			"message": err2.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, successMessage)
}

func (u *userRestHandler) UpdateUserData(c *gin.Context) {
	var updateUserDataRequest dto.UpdateUserDataRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&updateUserDataRequest)
	} else {
		err = c.ShouldBind(&updateUserDataRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	// ! TODO: Update error but data updated
	response, err2 := u.userService.UpdateUserData(userData.ID, &updateUserDataRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (u *userRestHandler) DeleteUser(c *gin.Context) {
	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	response, err2 := u.userService.DeleteUser(userData.ID)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
