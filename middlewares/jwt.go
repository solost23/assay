package middlewares

import (
	"assay/dao"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/middleware"
	"assay/infra/response"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, code, err := GetUserFromToken(c)
		if err != nil {
			response.Error(c, code, err)
			return
		}

		c.Set("user", user)
		c.Set("userId", user.ID)
		c.Next()
	}
}

// GetUserFromToken 从 token 中获取用户信息
func GetUserFromToken(c *gin.Context) (*dao.User, int, error) {
	token := c.Request.Header.Get("token")
	var user *dao.User
	if token == "" {
		return user, constant.UnauthorizedCode, errors.New("请登录")
	}

	j := middleware.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if errors.Is(err, middleware.ErrTokenExpired) {
			return user, constant.UnauthorizedCode, errors.New("授权已过期")
		}
		return user, constant.UnauthorizedCode, errors.New("无效 token")
	}

	rdb := global.RDB

	jwtToken, err := rdb.Get(c, claims.Device+strconv.Itoa(int(claims.ID))).Result()
	if err != nil {
		return user, constant.UnauthorizedCode, errors.New("无效 token")
	}
	if jwtToken != token {
		return user, constant.UnauthorizedCode, errors.New("无效 token")
	}
	jsonUser, err := rdb.Get(c, claims.Device+token).Result()
	if err != nil {
		return user, constant.UnauthorizedCode, errors.New("无效 token")
	}

	user = &dao.User{}
	err = json.Unmarshal([]byte(jsonUser), user)
	if err != nil {
		return user, constant.UnauthorizedCode, err
	}

	return user, 0, nil
}
