package controllers

import (
	"example/funtion/initializers"
	"example/funtion/models"
	"net/http"
	"os"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating user account"})
			return
		}
		
		user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: string(hash), Phone: body.Phone, Age: uint64(body.Age)}
		newUserAccount := initializers.DB.Create(&user)

		if newUserAccount.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating user account"})
			return
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
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error Message":"Inavalid email address or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {		
			c.JSON(http.StatusNotFound, gin.H{"Error Message":"Inavalid email address or password"})
			return
		
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": user.ID,
		"expires": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"failed to create token"})
		return	
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString,3600 * 24* 30, "", "", false, true)
	
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})



}