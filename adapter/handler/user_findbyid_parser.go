// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_findbyid_parser.go
UserFindByIdParser
*/

package handler

import (
	//"fmt"

	//	"encoding/json"

	"anemone/model"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/07/17 UserFindByIdParser 新規作成

type userFindByIdParser struct {
	model model.User
}

func NewUserFindByIdParser() *userFindByIdParser {
	return &UserCreateParser{model.User{}}
}

func (p *userFindByIdParser) TryParseFormat(data string) error {
	fmt.Printf("%#v", p.model)
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
