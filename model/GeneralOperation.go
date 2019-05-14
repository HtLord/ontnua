package model

import "github.com/gin-gonic/gin"

func Wrapper(obj *interface{}, c gin.Context) error {
	err := c.ShouldBind(obj)
	if err != nil {
		return err
	}
	return nil
}