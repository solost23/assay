package services

import (
	"assay/constants"
	"assay/dao"
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/middleware"
	"assay/infra/response"
	"assay/infra/util"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type LoginService struct{}

func (*LoginService) Login(c *gin.Context, params *forms.LoginForm) {
	var (
		sqlUser *dao.User
		err     error
	)

	db := global.DB

	switch params.Type {
	case 0:
		sqlUser, err = dao.GWhereFirstSelect[dao.User](db, "*", "username = ? AND password = ?", params.Username, util.NewMd5(params.Password, constant.Secret))
	case 1:
		// 对比redis中的短信验证码，然后去数据库拿到用户信息
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.BadRequestCode, errors.New("用户名或密码错误"))
		return
	}

	token, err := loginInitUserToken(c, sqlUser, params.PlatForm)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, &forms.Login{
		Token: token,
		User:  *sqlUser,
	})
}

func loginInitUserToken(c *gin.Context, user *dao.User, platform string) (string, error) {
	var redisPrefix string
	switch platform {
	case "assay":
		redisPrefix = constants.AssayRedisPrefix
	}
	token, err := loginGetToken(user.ID, redisPrefix)
	if err != nil {
		return "", err
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	rdb := global.RDB
	key := redisPrefix + strconv.Itoa(int(user.ID))
	oldToken, _ := rdb.Get(c, key).Result()

	rdb.Del(c, redisPrefix+oldToken)
	rdb.Set(c, key, token, time.Duration(global.ServerConfig.JWT.Duration)*time.Second)
	rdb.Set(c, redisPrefix+token, userJson, time.Duration(global.ServerConfig.JWT.Duration)*time.Second)

	return token, nil
}

func loginGetToken(userId uint, redisPrefix string) (string, error) {
	claims := constant.CustomClaims{
		ID:     userId,
		Device: redisPrefix,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.ServerConfig.JWT.Duration) * time.Second)),
			Issuer:    "assay",
		},
	}

	j := middleware.NewJWT()
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
