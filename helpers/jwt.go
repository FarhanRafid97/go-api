package helpers

import (
	"github.com/golang-jwt/jwt/v4"
)

func Sign(Data map[string]interface{}, SecrePublicKeyEnvName string) (string, error) {

	expiredAt := 60 * 60 * 1000 * 7

	jwtSecretKey := LoadEnv(SecrePublicKeyEnvName)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {

		return accessToken, err
	}

	return accessToken, nil
}
