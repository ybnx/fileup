package util

import (
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/jordan-wright/email"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

const (
	VeCodeNS  = "VeCodeNS"
	VersionNS = "VersionNS"
)

type Claims struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Version int64  `json:"version"`
	jwt.StandardClaims
}

func InitDB(dataSourece string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSourece))
	if err != nil {
		return nil
	}
	return db
}

func InitRDB(address string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
}

func GenerateToken(id, version int64, name, email, secret string, expire int64) (string, string, error) {
	accessClaims := Claims{
		Id:      id,
		Name:    name,
		Email:   email,
		Version: version, // 保证token只能使用一次
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expire) * time.Second).Unix(), // ExpiresAt就可以保证token只能使用一次，不需要version
			Issuer:    "token",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}
	refreshClaims := Claims{
		Id:      id,
		Name:    name,
		Email:   email,
		Version: version,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Duration(expire) * time.Second).Unix(),
			Issuer:    "token",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}
	return accessStr, refreshStr, nil
}

//func ParseToken(accessToken, refreshToken, jwtSecret, version string) bool {
//	claims := Claims{}
//	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
//		return []byte(jwtSecret), nil
//	})
//	num, err := strconv.Atoi(version)
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid && claims.Version == num {
//		return true
//	}
//	return false
//}

func GenerateCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < 6; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendCode(username, password, to, code string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("fileup <%s>", username)
	e.To = []string{to}
	e.Subject = "verify code"
	e.HTML = []byte("your verify code is: <b>" + code + "</b>")
	return e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", username, password, "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}

func Hash(password string) string { // TODO 加盐
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
