package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func BlackListToken(token string){
	blackList[token] = true
}

func GenerateAccess(userID uint, role string) (string,error) {
	secret_key := []byte(os.Getenv("SECRET_KEY"))
	claims := CustomClaims{
		UserID:           userID,
		Role:             role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15*time.Minute)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString(secret_key)
}

func GenerateRefresh(userID uint)(string, error){
	secret_key := []byte(os.Getenv("SECRET_REFRESH_KEY"))
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*7*time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret_key)
}


func VerifyToken(tokenString string, keyName string)(*CustomClaims, error){
	token,err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, 
		func(token *jwt.Token)(interface{},error){
			if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, errors.New("Sai thuật toán")
			}
			return []byte(os.Getenv(keyName)), nil
		},
	)
	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid{
		return claims, nil
	}
	return nil, errors.New("Token không hợp lệ")
}