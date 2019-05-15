package model

import (
	"time"
)

type News struct {
	Ban string `json:"ban" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	CreateTime time.Time `json:"createTime"`
}

