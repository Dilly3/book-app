package jwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"

	"os"
	"time"
)

type UserClaims struct {
	jwt.StandardClaims
	Email     string `json:"email"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	ExpireAt  int64  `json:"expires_at"`
	SessionID int64  `json:"session_id"`
}

func (u *UserClaims) Valid() error {
	if u.ExpireAt < time.Now().Unix() {
		return fmt.Errorf("%s", "token expired")

	}
	if u.SessionID == 0 {
		return fmt.Errorf("%s", "invalid session id")
	}
	return nil

}
func GenToken(u *UserClaims) (*string, error) {
	var JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" || len(JWT_SECRET) < 1 {
		JWT_SECRET = "TellSomeOneOfWhatYouKnowOrYouWontKnowAgain!!!!!!!!"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	ss, err := token.SignedString([]byte(JWT_SECRET))
	fmt.Printf("%v %v", ss, err)
	return &ss, nil
}

func ParseToken(tokenString string) (*UserClaims, error) {
	var JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" || len(JWT_SECRET) < 1 {
		JWT_SECRET = "TellSomeOneOfWhatYouKnowOrYouWontKnowAgain!!!!!!!!"
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	claims := token.Claims.(*UserClaims)
	if claims.Valid() != nil {
		return nil, fmt.Errorf("%s", "invalid token")
	}

	return claims, err

}
