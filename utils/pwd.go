package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"hmshop/global"
	"log"
	"time"
)

// HashPwd加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// CheckPwd验证密码 hashedPwd hash之后的密码  plainPwd输入的密码
func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// CustomPayload 自定义载荷继承原有接口并附带自己的字段
type CustomPayload struct {
	UserId     uint64 //签发人id
	GrantScope string //签发人
	jwt.RegisteredClaims
}

// GenerateToken 生成Token uid 用户id subject 签发对象  secret 加盐
func GenerateToken(c *gin.Context, uid uint64, subject string, secret string) (string, error) {
	claim := CustomPayload{
		UserId:     uid,
		GrantScope: subject,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   //签发者
			Subject:   subject,                                         //签发对象
			Audience:  jwt.ClaimStrings{"PC", "Wechat_Program"},        //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))

	exp := claim.RegisteredClaims.ExpiresAt
	now := claim.RegisteredClaims.IssuedAt

	diff := exp.Time.Sub(now.Time)
	err = LoginToken(c, token, diff, uid)
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseToken(token string, secret string) (*CustomPayload, error) {
	// 解析token
	parseToken, err := jwt.ParseWithClaims(token, &CustomPayload{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parseToken.Claims.(*CustomPayload); ok && parseToken.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func LoginToken(c *gin.Context, token string, diff time.Duration, uid uint64) error {
	//将注销用户的token放入redis中
	err := global.Redis.Set(c, "token:"+token, uid, diff).Err()
	return err
}

func LogoutToken(c *gin.Context, token string) error {
	// 将 Redis 中的 token 设置为过期
	err := global.Redis.Expire(c, "token:"+token, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
