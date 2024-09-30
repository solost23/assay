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
	"assay/infra/util/utf8togbk"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/tarm/serial"
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
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, constant.InternalServerErrorCode, err)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, constant.BadRequestCode, errors.New("用户名或密码错误"))
			return
		}
	case 1:
		// 对比redis中的短信验证码，然后去数据库拿到用户信息
		rdb := global.RDB
		code, err := rdb.Get(c, constants.AssayVerifyCodeRedisPrefix+params.Phone).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			response.Error(c, constant.InternalServerErrorCode, err)
			return
		}
		if errors.Is(err, redis.Nil) {
			response.Error(c, constant.InternalServerErrorCode, errors.New("验证码已过期"))
			return
		}
		if code != params.Code {
			response.Error(c, constant.InternalServerErrorCode, errors.New("验证码错误"))
			return
		}
		sqlUser, err = dao.GWhereFirstSelect[dao.User](db, "*", "phone = ?", params.Phone)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, constant.InternalServerErrorCode, err)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, constant.BadRequestCode, errors.New("手机号不存在"))
			return
		}
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
		redisPrefix = constants.AssayLoginRedisPrefix
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

func (*LoginService) Code(c *gin.Context, params *forms.LoginGetCodeForm) {
	catConfig := global.ServerConfig.Cat

	s, err := serial.OpenPort(&serial.Config{
		Name:        catConfig.Name,
		Baud:        catConfig.Baud,
		ReadTimeout: time.Duration(catConfig.ReadTimeout) * time.Millisecond,
		Size:        catConfig.Size,
		Parity:      serial.Parity(catConfig.Parity),
		StopBits:    serial.StopBits(catConfig.StopBits),
	})
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	defer s.Close()

	code := util.GenerateRandCode()
	data := fmt.Sprintf("#%s#%s#", params.Phone, fmt.Sprintf("【北方稀土】您的验证码是%s，在15分钟内有效。如非本人操作请忽略本短信。", code))
	b, err := utf8togbk.UTF8ToGBK([]byte(data))
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if _, err = s.Write(b); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	rdb := global.RDB
	if err := rdb.Set(c, constants.AssayVerifyCodeRedisPrefix+params.Phone, code, time.Duration(15)*time.Minute).Err(); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, "success")
}
