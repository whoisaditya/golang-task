package controllers

import (
	"net/http"
	"os"
	"time"

	"golang-task/initializers"
	"golang-task/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get email/password from request body
	var body struct {
		Email string 
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return	
	}


	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	// create user
	user := models.User{ Email: body.Email, Password: string(hash) }
	result := initializers.DB.Create(&user)
	
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating user",
		})
		return
	}


	// respond
	c.JSON(http.StatusOK, gin.H{ 
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	// get email/password from request body
	var body struct {
		Email string 
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return
	}

	// find user by email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}


	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),	
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error generating jwt",
		})
		return
	}

	// respond with jwt
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
	})
}

func AddDetails(c *gin.Context) {
	// get email/password from request body
	var body struct {
		Email string
		Name string
		Phone string
		Age int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return	
	}

	// find user by email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	// update user
	user.Name = body.Name
	user.Phone = body.Phone
	user.Age = body.Age
	initializers.DB.Save(&user)

	// respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Details added successfully",
	})	
}