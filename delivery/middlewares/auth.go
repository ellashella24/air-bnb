package middlewares

import (
	"air-bnb/constants"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Auth interface {
	GenerateToken(userID int) (string, error)
	ExtractTokenUserID(e echo.Context) int
	ExtractTokenEmail(e echo.Context) string
	IsAdmin(next echo.HandlerFunc) error
}

type jwtService struct {
}

func NewAuth() *jwtService {
	return &jwtService{}
}

func (a *jwtService) GenerateToken(userID int, email, role string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["email"] = email
	claim["role"] = role
	claim["exp"] = time.Now().Add(time.Hour * 12).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(constants.SecretKey))
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (a *jwtService) ExtractTokenUserID(c echo.Context) int {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)

	if !strings.Contains(authHeader, "Bearer") {
		return 0
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(constants.SecretKey), nil
	})

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return 0
	}

	userID := int(claim["user_id"].(float64))

	return userID

}

func (a *jwtService) ExtractTokenEmail(c echo.Context) string {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)

	if !strings.Contains(authHeader, "Bearer") {
		return ""
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(constants.SecretKey), nil
	})

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return ""
	}

	email := claim["email"].(string)

	return email

}

func (a *jwtService) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(echo.HeaderAuthorization)

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(constants.SecretKey), nil
		})

		claim, _ := token.Claims.(jwt.MapClaims)

		role := claim["role"]
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "not admin",
			})
		}

		return next(c)
	}
}

// func (a *jwtService) ExtractTokenUserID(c echo.Context) int {
// 	user := c.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		userId := int(claims["userID"].(float64))
// 		return userId
// 	}
// 	return 0
// }

// func (a *jwtService) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(jwt.MapClaims)
// 		roles := claims["roles"].(string)
// 		if roles != "admin" {
// 			return echo.ErrUnauthorized
// 		}
// 		return next(c)
// 	}
// }
