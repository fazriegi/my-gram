package helper

import "github.com/dgrijalva/jwt-go"

var secret = "test"

func GenerateToken(id int, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secret))

	return signedToken, err
}
