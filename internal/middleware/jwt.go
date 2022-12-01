package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const jwtKey = "lafnlksnfdklasnfa"

// ContextKeySub is where JWT for user id is stored in r.Context
var ContextKeySub contextKey = "sub"

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

func (c contextKey) String() string {
	return string(c)
}

func GenerateToken(userID string) (string, error) {
	mySigningKey := []byte(jwtKey)

	// Create the claims
	claims := MyCustomClaims{
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// AuthorizeJWT verify token and add userID to the request context
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to process request"})
			c.Abort()
			return
		}

		var claims MyCustomClaims
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			//Make sure that the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtKey), nil
		})

		//Token is invalid, maybe not signed on this server
		if !token.Valid {
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
				c.Abort()
				return
			case errors.Is(err, jwt.ErrTokenMalformed):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "That's not even a token"})
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			c.Abort()
			return
		}

		if len(claims.Subject) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			c.Abort()
			return
		}

		c.Set(ContextKeySub.String(), claims.Subject)
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	bearToken := c.GetHeader("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		return "", errors.New("[extractToken] Token is not valid")
	}

	return strArr[1], nil
}
