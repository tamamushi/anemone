// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_update_parser.go
UserUpdateParser
*/

package handler

import (
	"anemone/adapter/gateway"
	"anemone/model"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/06/27 UserUpdateParser 新規作成

type UserUpdateParser struct {
	model *model.User
	gateway.IFormatParser
	gateway.IInput
	gateway.IOutput
}

func NewUserUpdateParser(
	parser gateway.IFormatParser,
	input gateway.IInput,
	output gateway.IOutput,
) *UserUpdateParser {
	return &UserUpdateParser{
		new(model.User), parser, input, output,
	}
}

func (p *UserUpdateParser) GetModel() *model.User {
	return p.model
}

func (p *UserUpdateParser) SetParser(interface{}) {
	return
}
