// +build -user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_create_parser.go
UserCreateParser
*/

package handler

import (
	"anemone/model"
	"encoding/json"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/06/27 UserCreateParser 新規作成

type UserCreateParser struct {
	model model.User
}

func NewUserCreateParser() *UserCreateParser {
	return &UserCreateParser{model.User{}}
}

func (p *UserCreateParser) GetModel() *model.User {
	return &p.model
}

func (p *UserCreateParser) TryParseFormat(data string) error {
	if err := json.Unmarshal([]byte(data), &p.model); err != nil {
		return err
	}
	return nil
}

func (p *UserCreateParser) Input(data string) interface{} {
	if err := json.Unmarshal([]byte(data), &p.model); err != nil {
		return err
	}
	return p.model
}
