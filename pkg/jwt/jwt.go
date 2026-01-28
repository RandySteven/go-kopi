package jwt_client

import (
	"os"
	"time"

	"github.com/RandySteven/go-kopi/entities/models"
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type (
	JWTAccessClaim struct {
		UserID   uint64
		Username string
		RoleID   []uint64
		IsVerify bool
		jwt.RegisteredClaims
	}

	JWTRefreshClaim struct {
		UserID uint64
		Email  string
		jwt.RegisteredClaims
	}
)

func GenerateTokens(user *models.User) (string, string) {

	access := &JWTAccessClaim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Applications",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	//for _, role := range roleUser {
	//	access.RoleID = append(access.RoleID, role.RoleID)
	//}
	tokenAccessAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, access)
	accessToken, err := tokenAccessAlgo.SignedString(JwtKey)
	if err != nil {
		return "", ""
	}

	refresh := &JWTRefreshClaim{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Applications",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
		},
	}
	tokenRefreshAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh)
	refreshToken, err := tokenRefreshAlgo.SignedString(JwtKey)
	if err != nil {
		return "", ""
	}

	return accessToken, refreshToken
}
