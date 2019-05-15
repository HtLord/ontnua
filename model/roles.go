package model

import (
	"time"
)

type ADMLevel int

const (
	ROOT_ADM ADMLevel = 0
)

type Member struct {
	Name string `json:"name" form:"name" xml:"name" binding:"required"`
	Email string `json:"email" form:"email" xml:"email" binding:"required"`
	Secret string `json:"secret" form:"secret" xml:"secret" binding:"required"`
	CreateTime time.Time `json:"createTime" form:"createTime" xml:"createTime"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" xml:"updateTime"`
}

type ADM struct {
	Member
	Level ADMLevel
}

type Store struct {
	BAN string `msgpack:"ban" json:"ban" form:"ban" xml:"ban" binding:"required"`
	Name string `msgpack:"name" json:"name" form:"name" xml:"name" binding:"required"`
	ADMs []ADM `msgpack:"adms" json:"adms" form:"adms" xml:"adms" binding:"required"`
	Members []Member `msgpack:"members" json:"members" form:"members" xml:"members" binding:"required"`
}
