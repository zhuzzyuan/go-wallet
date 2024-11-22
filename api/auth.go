package api

import (
	"go-wallet/api/models"
	"go-wallet/config"
	"go-wallet/render"
	"go-wallet/util/log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetToken(c *gin.Context) {
	var req models.GetTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, render.BindError(err))
		return
	}

	result := generateNewToken(req.Email)
	c.JSON(200, result)
}

func generateNewToken(email string) models.Response {
	tokenConfig := config.GetTokenConfig()

	tokenClaims := models.CustomTokenClaims{
		Email: email,
		Type:  "token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenConfig.AccessTokenExpiry).Unix(),
		},
	}

	refreshTokenClaims := models.CustomTokenClaims{
		Email: email,
		Type:  "refresh_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenConfig.RefreshTokenExpiry).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenString, err := token.SignedString([]byte(tokenConfig.JWTSecret))
	if err != nil {
		log.Error(err)
		return render.Error(err)
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(tokenConfig.JWTSecret))
	if err != nil {
		log.Error(err)
		return render.Error(err)
	}

	tokenResp := &models.TokenResponse{
		TokenType:    "Bearer",
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    tokenConfig.AccessTokenExpiry.Milliseconds() / 1000,
	}

	return render.Success(tokenResp)
}
