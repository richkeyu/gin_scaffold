package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gateway/pkg/client"

	"github.com/richkeyu/gocommons/plog"
	jsoniter "github.com/json-iterator/go"

	"gateway/pkg/codes"
	"gateway/pkg/jwt"

	"github.com/richkeyu/gocommons/util/app"

	jwt2 "github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

const (
	HeaderKeyUser = "x-user-data"
)

type merchantUser struct {
	Uid        int64  `json:"uid"`
	Email      string `json:"email"`
	MerchantId int64  `json:"merchant_id"`
	Name       string `json:"name"`
}

// AuthMiddleWare
// 登录校验
// 根据service参数来确定用户是企业用户还是商城用户
func initAuthMiddleWare(service string, auth bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var claims jwt2.MapClaims
		authorization := ctx.GetHeader("Authorization")
		if authorization != "" && strings.HasPrefix(authorization, "Bearer ") {
			var err error
			bearer := strings.TrimPrefix(authorization, "Bearer ")
			claims, err = jwt.ParseToken(bearer, service)
			if err != nil && auth {
				// 验证失败 如果存在商城cookie尝试删除旧cookie
				cookie, err := ctx.Request.Cookie("site-token")
				if err == nil && len(cookie.Value) > 0 {
					http.SetCookie(ctx.Writer, &http.Cookie{
						Name:     "site-token",
						Value:    url.QueryEscape(""),
						MaxAge:   0,
						Path:     "/",
						Domain:   "",
						Secure:   true,
						HttpOnly: false,
						Expires:  time.Now(),
					})
				}

				ctx.Abort()
				app.Error(ctx, http.StatusOK, codes.ErrorPermissionDenied)
				return
			}

			claimsJson, _ := json.Marshal(claims["data"])
			ctx.Request.Header.Set(HeaderKeyUser, string(claimsJson))

			if service == jwt.Merchant {
				// 查询用户信息
				var info merchantUser
				err = jsoniter.Unmarshal(claimsJson, &info)
				if err != nil {
					plog.Error(ctx, "AuthMiddleWare Unmarshal fail: %s", err)
					ctx.Abort()
					app.Error(ctx, http.StatusOK, codes.ErrorPermissionDenied)
					return
				}
				user, err := client.NewBaseCli().GetUserByIdWithCache(ctx, info.Uid)
				if err != nil {
					plog.Error(ctx, "AuthMiddleWare get user fail: %s", err)
				} else {
					info.MerchantId = user.MerchantID
					info.Name = user.Name
					data, err := jsoniter.MarshalToString(info)
					if err != nil {
						plog.Error(nil, "AuthMiddleWare get user fail: %s", err)
					} else {
						ctx.Request.Header.Set(HeaderKeyUser, data)
					}
				}
			}
			ctx.Next()
		} else {
			if auth {
				ctx.Abort()
				app.Error(ctx, http.StatusOK, codes.ErrorPermissionDenied)
			}

			return
		}

	}
}

func MerchantAuthMiddleWare() gin.HandlerFunc {
	return initAuthMiddleWare(jwt.Merchant, true)
}

func StoreAuthMiddleWare() gin.HandlerFunc {
	return initAuthMiddleWare(jwt.Store, true)
}

func StoreWithOutAuthMiddleWare() gin.HandlerFunc {
	return initAuthMiddleWare(jwt.Store, false)
}
