package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/golang-jwt/jwt/v4"
	"github.com/greper/ulogin/internal/model/do"
	"github.com/greper/ulogin/internal/model/entity"
	"github.com/greper/ulogin/internal/service/setting"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	UserService *UserService
}

var AuthSvc = &AuthService{
	UserService: UserSvc,
}

type LoginByPasswordIn struct {
	Username string
	Password string
	UserType int64
}

type JwtClaims struct {
	UserId   int64
	RoleIds  []int64
	UserType int64 // 1 = admin,2 = user
	jwt.RegisteredClaims
}

type JwtToken struct {
	AccessToken string `json:"accessToken"`
	Expires     int64  `json:"expires"`
}

type UserInfo struct {
	Id       int64  `json:"id" dc:"用户Id"`
	Username string `json:"username" dc:"用户名"`
	Avatar   string `json:"avatar" dc:"头像"`
	NickName string `json:"nickName" dc:"昵称"`
}

type LoginSuccessOut struct {
	UserInfo   *UserInfo
	AccessInfo *JwtToken
}

func (s *AuthService) LoginByPassword(ctx context.Context, in *LoginByPasswordIn) (out *LoginSuccessOut, err error) {

	found, err := UserSvc.FindOne(ctx, &do.User{Username: in.Username, UserType: in.UserType})
	if err != nil {
		return nil, err
	}
	if found == nil {
		return nil, gerror.New("用户名或密码错误")
	}
	verified := s.CompareHashAndPassword(ctx, found.Password, in.Password)
	if !verified {
		return nil, gerror.New("用户名或密码错误")
	}
	token, err := s.BuildLoginReply(ctx, found)
	if err != nil {
		return nil, err
	}

	return &LoginSuccessOut{
		UserInfo: &UserInfo{
			Id:       found.Id,
			Username: found.Username,
			Avatar:   found.Avatar,
			NickName: found.NickName,
		},
		AccessInfo: token,
	}, nil
}

func (s *AuthService) CompareHashAndPassword(ctx context.Context, encrypted string, userInput string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(userInput)); err != nil {
		glog.Errorf(ctx, "密码校验失败, %+v", err)
		//password, err := s.EncryptPassword(userInput)
		//if err != nil {
		//	return false
		//}
		//glog.Infof(ctx, "密码：%s", password)
		return false
	}
	return true
}

func (s *AuthService) EncryptPassword(password string) (string, error) {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil { // 加密密码
		return "", err
	} else {
		return string(byTes), nil
	}
}

func (s *AuthService) BuildLoginReply(ctx context.Context, user *entity.User) (out *JwtToken, err error) {

	jwtSetting := &setting.JwtSetting{}
	err = SysSettingSvc.GetSetting(ctx, jwtSetting)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	expireAt := now.Add(time.Duration(jwtSetting.TokenExpires) * time.Hour)

	claims := JwtClaims{
		UserId:   user.Id,
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now.Add(-time.Duration(1000) * time.Second)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(expireAt),                                    // 过期时间 7天  配置文件
			Issuer:    "greper@ulogin",                                                 // 签名的发行者
		},
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	accessToken, err := token.SignedString([]byte(jwtSetting.JwtSecret))
	if err != nil {
		g.Log().Errorf(ctx, "生成JwtToken失败, %+v", err)
		return nil, err
	}
	return &JwtToken{
		AccessToken: accessToken,
		Expires:     expireAt.Unix(),
	}, nil
}
