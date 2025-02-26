package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"
	"github.com/greper/ulogin/internal/service"
	"github.com/greper/ulogin/internal/service/setting"
	"strings"
)

func JwtAuthMiddleware(r *ghttp.Request) {

	if strings.HasPrefix(r.RequestURI, "/api/admin/auth") {
		r.Middleware.Next()
		return
	}
	// Get token from Authorization header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "No token provided"))
		return
	}

	// Remove 'Bearer ' prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	jwtSetting := &setting.JwtSetting{}
	err := service.SysSettingSvc.GetSetting(r.Context(), jwtSetting)
	if err != nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "登录配置错误"))
	}
	JwtSecretKey := jwtSetting.JwtSecret
	// Parse and validate the token
	claims := &service.JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "Invalid token"))
		return
	}

	r.SetCtxVar("LoginUser", claims.LoginUser)
	r.Middleware.Next()
}
