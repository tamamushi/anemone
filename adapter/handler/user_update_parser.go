// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_update_parser.go
UserUpdateParser
*/

package handler

import (
	"anemone/model"
	"encoding/json"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/06/27 UserUpdateParser 新規作成

type UserUpdateParser struct {
	model model.User
}

func NewUserUpdateParser() *UserUpdateParser {
	return &UserUpdateParser{model.User{}}
}

func (p *UserUpdateParser) TryParseFormat(data string) error {
	if err := json.Unmarshal([]byte(data), &p.model); err != nil {
		return err
	}
	return nil
}

func (p *UserUpdateParser) Input(data string) interface{} {
	if err := json.Unmarshal([]byte(data), &p.model); err != nil {
		return err
	}
	return p.model
}
