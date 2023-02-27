package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/richkeyu/gocommons/config"
)

const (
	Merchant       = "merchant"
	Store          = "store"
	SECRETMERCHANT = "+matthg$w)x_6#ax0uug(jdy0)@^9(a)%w6nr$bmbcsac-n8x4"
)

type MerchantClaims struct {
	jwt.StandardClaims
	Data MerchantDataClaims `json:"data"`
}

type CustomerClaims struct {
	jwt.StandardClaims
	Data CustomerDataClaims `json:"data"`
}

type MerchantDataClaims struct {
	Uid   int    `json:"uid"`
	Email string `json:"email"`
}

type CustomerDataClaims struct {
	Uid       int    `json:"uid"`
	Email     string `json:"email"`
	LoginType int    `json:"login_type"`
}

// ParseToken
// 解析Token
func ParseToken(tokenString string, service string) (jwt.MapClaims, error) {
	secret := getSecretKey(service)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}

// CreateToken
// 生成Token：
// SecretKey 是一个 const 常量
func CreateToken(merchantDataClaims MerchantDataClaims) (tokenString string, err error) {
	now := time.Now().Unix()
	claims := &MerchantClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(now)),
			IssuedAt:  jwt.NewTime(float64(now + 7*24*3600)),
			Issuer:    "richkey",
		},
		Data: merchantDataClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(SECRETMERCHANT))
	return

}

type Conf struct {
	SecretKey       string `yaml:"secret_key"`
	ExpiredDuration int64  `yaml:"expired_duration"`
	Issuer          string `yaml:"issuer"`
}

func getConf(service string) Conf {
	var err error
	var confMap map[string]Conf

	if err = config.Load("jwt", &confMap); err != nil {
		panic(fmt.Sprintf("load jwt config failed, err: %v", err))
	}
	for key, value := range confMap {
		if key == service {
			return value
		}
	}
	panic(fmt.Sprintf("load jwt config failed, err: conf for #{service} not found"))
}

func getSecretKey(service string) string {
	conf := getConf(service)
	return conf.SecretKey
}
