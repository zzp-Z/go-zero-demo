package logic

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Claims 定义Claims
type Claims struct {
	UserId uint64 `json:"userId"`
	jwt.RegisteredClaims
}
type Tools struct {
	jwtKey         []byte
	expirationTime time.Time
}

func NewTools() *Tools {
	return &Tools{
		jwtKey:         []byte("my_secret_key"),
		expirationTime: time.Now().Add(2 * time.Minute),
	}
}

// HashPassword 生成哈希密码
func (t *Tools) HashPassword(password string) (hashedPassword string, err error) {
	byteHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword = string(byteHashedPassword)
	return
}

// CheckPassword 验证密码
func (t *Tools) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GenerateJwtToken 生成JwtToken
func (t *Tools) GenerateJwtToken(userId uint64) (tokenString string, err error) {
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t.expirationTime),
		},
	}
	// 创建JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名JWT
	tokenString, err = token.SignedString(t.jwtKey)
	if err != nil {
		return tokenString, &AppError{
			Code:    "TS0643",
			Message: "JWT签名失败",
		}
	}
	return
}

// ParseJwtToken 解析JwtToken
func (t *Tools) ParseJwtToken(tokenString string) (userId uint64, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return t.jwtKey, nil
	})
	if err != nil {
		return 0, &AppError{
			Code:    "TS0644",
			Message: "JWT解析失败",
		}
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, &AppError{
			Code:    "TS0645",
			Message: "JWT验证失败",
		}
	}
	return claims.UserId, nil
}

// GetNowTime 获取当前时间
func (t *Tools) GetNowTime() sql.NullTime {
	return sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}
