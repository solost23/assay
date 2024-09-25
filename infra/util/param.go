package util

import (
	"github.com/gin-gonic/gin"
)

func DefaultGetValidParams(c *gin.Context, params any) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	return nil
}

func GetValidUriParams(c *gin.Context, params any) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	return nil
}
