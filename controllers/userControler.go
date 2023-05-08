package controllers

import (
	"example/funtion/initializers"
	"example/funtion/models"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct{
		FirstName string 
		LastName string
		Email string
		Password string
		Age uint16
		Phone string
	}
	 
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error Message":"KIndly fill in all the details to sign up"})
			return
		}
		if len(body.Phone) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid phone number"})
        return 
		}
		for _, r := range body.Phone {
			if !unicode.IsDigit(r) {
				c.JSON(http.StatusBadRequest, gin.H{"message": "phone number can only contain numeric characters"})
				
				return 
			}
		}
		user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: body.Password, Phone: body.Phone, Age: uint64(body.Age)}
		newUserAccount := initializers.DB.Create(&user)

		if newUserAccount.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating user account"})
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully for "})
	 

}
func SignIn (c *gin.Context) {
	var body struct {
		Email string
		Password string
	}
	if c.Bind(body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meesage": "Kindly fill in a username and password",
		})
	}
	

}