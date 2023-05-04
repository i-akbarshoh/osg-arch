package jwt

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gookit/config/v2"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Tokens struct {
	Access    string
	AccExpire int64
	Refresh   string
}

// GenerateNewTokens func for generate a new Access & Refresh token
func GenerateNewTokens(id string, credentials map[string]string) (*Tokens, error) {
	accessToken, expire, err := generateNewAccessToken(id, credentials)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:    accessToken,
		Refresh:   refreshToken,
		AccExpire: expire,
	}, nil
}

func generateNewAccessToken(id string, credentials map[string]string) (string, int64, error) {
	claims := jwt.MapClaims{}

	claims["id"] = id
	claims["role"] = credentials["role"]

	if config.Data()["environment"].(string) == "develop" {
		claims["expires"] = time.Now().Add(time.Hour * 24 * 31).Unix()
	} else {
		claims["expires"] = time.Now().Add(time.Minute * 2 * time.Duration(config.Data()["access_expire"].(int))).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Data()["signing_key"].(string)))
	if err != nil {
		return "", 0, err
	}

	return t, claims["expires"].(int64), nil
}

func generateNewRefreshToken() (string, error) {
	sha256Hash := sha256.New()

	refresh := config.Data()["refresh_key"].(string) + time.Now().String()

	_, err := sha256Hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}

	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(config.Data()["refresh_expire"].(int))).Unix())

	t := hex.EncodeToString(sha256Hash.Sum(nil)) + "." + expireTime

	return t, nil
}
