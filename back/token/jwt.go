package token

import (
	"fmt"
	"log"
	"temsys"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Tokenizer is an implementation of a user Tokenizer using JSON Web Tokens.
type Tokenizer struct {
	signKey string
}

// New JWT Tokenizer
func New(signKey string) Tokenizer {
	return Tokenizer{signKey}
}

// Tokenize a user using JWT to create a Token
func (t Tokenizer) Tokenize(user temsys.User) temsys.Token {
	const hoursToExpire int = 5
	expiration := time.Now().Add(time.Hour * 5)
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["owner"] = user.Name
	token.Claims["expires"] = expiration.Unix()
	token.Claims["role"] = user.Role
	token.Header["alg"] = "HS256"
	tokenStr, err := token.SignedString([]byte(t.signKey))
	if err != nil {
		log.Printf("Something wrong occurred while signing a JWT token for user: '%s': %s\n", user.Name, err)
		tokenStr = ""
	}
	return temsys.Token{
		Owner:   user.Name,
		Expires: expiration,
		Value:   tokenStr,
		Role:    user.Role,
	}
}

// Decode raw JWT token and stores the result into a Token struct.
// It can return an error if something wrong happens when decoding.
func (t Tokenizer) Decode(raw string) (temsys.Token, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) ([]byte, error) {
		if token.Header["alg"] != "HS256" {
			return nil, fmt.Errorf("Unknown alg")
		}
		return []byte(t.signKey), nil
	})
	if err != nil || !token.Valid {
		return temsys.Token{}, err
	}
	roleRaw := token.Claims["role"].(string)
	role := temsys.Role(roleRaw)
	return temsys.Token{
		Expires: time.Unix(int64(token.Claims["expires"].(float64)), 0),
		Owner:   token.Claims["owner"].(string),
		Value:   raw,
		Role:    role,
	}, nil
}
