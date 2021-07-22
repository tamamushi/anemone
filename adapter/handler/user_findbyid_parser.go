// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_findbyid_parser.go
UserFindByIdParser
*/

package handler

import (
	"anemone/adapter/gateway"
	"anemone/model"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/07/17 UserFindByIdParser 新規作成
// 2021/07/26 UserFindByIdParser IFormatParserとIOutputに対応させる

type UserFindByIdParser struct {
	model *model.User
	gateway.IFormatParser
	gateway.IOutput
}

func NewUserFindByIdParser(
	parser gateway.IFormatParser,
	output gateway.IOutput,
) *UserFindByIdParser {

	return &UserFindByIdParser{
		new(model.User),
		parser,
		output,
	}
}

func (p *UserFindByIdParser) SetParser(interface{}) {
	return
}
func (p *UserFindByIdParser) GetModel() *model.User {
	return p.model
}
