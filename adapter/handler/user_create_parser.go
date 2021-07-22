// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_create_parser.go
UserCreateParser
*/

package handler

import (
	"anemone/adapter/gateway"
	"anemone/model"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/06/27 UserCreateParser 新規作成
// 2021/07/26 UserCreateParser IFormatParserとIInput、IOutputに対応させる

type UserCreateParser struct {
	model *model.User
	gateway.IFormatParser
	gateway.IInput
	gateway.IOutput
}

func NewUserCreateParser(
	parser gateway.IFormatParser,
	input gateway.IInput,
	output gateway.IOutput,
) *UserCreateParser {
	return &UserCreateParser{
		new(model.User), parser, input, output,
	}
}

func (p *UserCreateParser) GetModel() *model.User {
	return p.model
}

func (p *UserCreateParser) SetParser(interface{}) {
	return
}
