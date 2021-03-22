package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main() {
	fmt.Printf("%T\n", time.Duration(1).Seconds())
	u := UserInfo{ID: 234, Username: "yida", StandClaims: jwt.StandardClaims{ExpiresAt: int64(time.Duration(2).Seconds())}}
	token, err := CreateToken(&u)
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", token)
	fmt.Println("---------------")
	// time.Sleep(time.Second * 1)
	userInfo, e := ParseToken(token)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%+v\n", userInfo)
}

type UserInfo struct {
	ID          uint64
	Username    string
	StandClaims jwt.StandardClaims
}

func CreateToken(user *UserInfo) (tokens string, err error) {
	// 自定义claim
	claim := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      user.StandClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err = token.SignedString([]byte("mySecretTest"))
	return
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("mySecretTest"), nil
	}
}

func ParseToken(tokens string) (user *UserInfo, err error) {
	user = &UserInfo{}
	token, err := jwt.Parse(tokens, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	fmt.Printf("token.Claims:%+v\n", token.Claims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	// 验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	fmt.Printf("claim:%+v\n", claim)

	user.ID = uint64(claim["id"].(float64))
	user.Username = claim["username"].(string)
	return
}
