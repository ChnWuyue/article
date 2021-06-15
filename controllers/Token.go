package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JWTClaim struct {
	userID string `json:"user_id"`
	jwt.StandardClaims
}

//signature
var (
	JWTSecret           = []byte("为什么不惜那样地努力呢？")
	TokenExpireDuration = time.Hour * 2
)

// @Title [Authorize needed] NewJwt
// @Description get JWT
func NewJWT(userName string) (string, error) {
	//创建一个jwt Claim
	j := JWTClaim{
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "wuyue",
		},
	}

	//指定加密方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, j)
	//加入签证
	return token.SignedString(JWTSecret)
}

// @Title [Authorize needed] ParseToken
// @Description  Parse and Check token
func ParseToken(tokenString string) (*JWTClaim, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	//检验token
	if claims, isOk := token.Claims.(*JWTClaim); isOk && token.Valid {
		tokenRefresh()
		return claims, nil
	}

	return nil, errors.New("invalid")

}

// @Title [Authorize needed] CheckJWT
// @Description Check JWT Middleware
func CheckJWT(c *gin.Context) {
	// Token放在Header的Authorization中

	authHead := c.Request.Header.Get("Authorization")
	if authHead == "" {
		c.JSON(http.StatusOK, gin.H{
			"error":   -2,
			"message": "Authorization is empty.",
		})
		c.Abort()

		return
	}

	mc, err := ParseToken(authHead)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   -3,
			"message": "无效的token",
		})

		c.Abort()
		return
	}
	//将请求的username信息保存到请求context中
	c.Set("userId", mc.userID)
	c.Next()
	return
}

// @Title [Authorize needed] TokenRefresh()
// @Description Check JWT Middleware
func tokenRefresh() {

}
