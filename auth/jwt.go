package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TokenS struct {
	User_id string `json:"user_id" binding:"required"`
}

func GenerateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return ([]byte(os.Getenv("SECRET_KEY"))), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	// 1️⃣ Call jwt.Parse(tokenString, callback)
	// 2️⃣ Extract Header, Payload, and Signature from JWT
	// 3️⃣ Detect signing method (e.g., HS256)
	// 4️⃣ Call the callback function → return SECRET_KEY
	// 5️⃣ Use SECRET_KEY to verify signature
	//    ✅ If valid → Continue to extract claims
	//    ❌ If invalid → Return error
	// 6️⃣ Return parsed token object or error

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("claims Invalid")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("claims Invalid")
	}
	return userID, nil
}

func HandleJWT(c *gin.Context) {
	var obj1 TokenS
	if err := c.ShouldBindJSON(&obj1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": err,
		})
		return
	}
	fmt.Println(obj1.User_id)
	token, err1 := GenerateJWT(obj1.User_id)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": err1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
