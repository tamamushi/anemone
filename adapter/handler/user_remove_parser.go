// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_remove_parser.go
UserRemoveParser
*/

package handler

import (
	"anemone/adapter/gateway"
	"anemone/model"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/07/26 UserRemoveParser 新規作成

type UserRemoveParser struct {
	model *model.User
	gateway.IFormatParser
}

func NewUserRemoveParser(
	parser gateway.IFormatParser,
) *UserRemoveParser {
	return &UserRemoveParser{
		new(model.User), parser,
	}
}

func (p *UserRemoveParser) GetModel() *model.User {
	return p.model
}

func (p *UserRemoveParser) SetParser(interface{}) {
	return
}
