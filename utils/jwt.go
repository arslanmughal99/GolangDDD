package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var (
	_ = godotenv.Load()

	key        = getKey()
	expiration = getExp()
)

type ValidJwt struct {
	Verified bool
	Username string
}

func CreateJwt(username string, verified bool) (*string, error) {

	now := time.Now()
	token := jwt.New()

	_ = token.Set("username", username)
	_ = token.Set("verified", verified)
	_ = token.Set(jwt.IssuedAtKey, now.Unix())
	_ = token.Set(jwt.ExpirationKey, now.Add(time.Hour*time.Duration(expiration)).Unix())

	rawToken, err := jwt.Sign(token, jwa.HS256, key)

	if err != nil {
		return nil, err
	}

	jwtToken := string(rawToken)

	return &jwtToken, nil
}

func ValidateJwt(token string) *ValidJwt {
	vt := new(ValidJwt)
	t, err := jwt.ParseString(token, jwt.WithVerify(jwa.HS256, key))
	if err != nil {
		return nil
	}

	verified, _ := t.Get("verified")
	username, _ := t.Get("username")

	vt.Verified = verified.(bool)
	vt.Username = fmt.Sprintf("%s", username)

	return vt
}

func getExp() int {
	exp, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP"))

	if err != nil {
		panic("Failed to convert env `ACCESS_TOKEN_EXP` to int")
	}

	return exp
}

func getKey() []byte {
	k := os.Getenv("JWT_AUTH_KEY")
	if k == "" {
		panic("No jwt signing key provided")
	}

	return []byte(k)
}
