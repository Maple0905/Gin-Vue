package controllers

import (
	"fmt"
	"gin-vue/models"
	"gin-vue/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messaage": "success", "data": u})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	us := models.GetUserByEmail(input.Email)
	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password
	token, err := models.LoginCheck(u.Email, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Password is incorrect"})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"id":          us.ID,
			"username":    us.Username,
			"email":       us.Email,
			"role":        models.GetRoleName(3),
			"accessToken": token,
		})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Passwrod string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Passwrod
	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration success!"})
}

func GetUserProfile(c *gin.Context) {

	fmt.Println("entered")
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": gin.H{"id": u.ID, "username": u.Username, "email": u.Email, "createdAt": "2022-11-08T14:22:27.000Z", "updatedAt": "2022-11-08T14:22:27.000Z", "roleId": u.Roleid}})

	// var input GetUserProfileInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// u := models.User{}
	// u.ID = int(input.id)

	// us, err := models.GetUserByID(input.id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// fmt.Println(us)

	// response := Body{"id": 1, "username": "admin", "email": "admin@gmail.com", "createdAt": "2022-11-08T14:22:27.000Z", "updatedAt": "2022-11-08T14:22:27.000Z", "roleId": 3}
	// c.JSON(http.StatusOK, gin.H{"profile": gin.H{"id": 1, "username": "admin", "email": "admin@gmail.com", "createdAt": "2022-11-08T14:22:27.000Z", "updatedAt": "2022-11-08T14:22:27.000Z", "roleId": 3}})
}
