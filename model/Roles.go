package model

import (
	"time"
)

type ADMLevel int

const (
	ROOT_ADM ADMLevel = 0
)

type Member struct {
	Name string `json: "name" binding: "required"`
	Email string `json: "email" binding: "required"`
	Secret string `json: "secret" binding: "required"`
	CreateTime time.Time `json: "createTime" binding: "required"`
	UpdateTime time.Time `json: "updateTime" binding: "required"`
}

type ADM struct {
	Member
	Level ADMLevel
}

type Store struct {
	BAN string `json: "ban" binding: "required"`
	ADMs []ADM `json: "adms" binding: "required"`
	Members []Member `json: "members" binding: "required"`
}
