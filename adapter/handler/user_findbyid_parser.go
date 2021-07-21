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
	"anemone/adapter/gateway"
	"anemone/model"
)

// TODO(koube):
//
// HISTORY(koube):
// 2021/07/17 UserFindByIdParser 新規作成

type UserFindByIdParser struct {
	model model.User
//	*gateway.IFormatParser
	*gateway.Output
}

func NewUserFindByIdParser(
//	parser *gateway.IFormatParser,
	output *gateway.Output,
) *UserFindByIdParser {

	return &UserFindByIdParser{
		model.User{},
//		parser,
		output,
	}
}

func (p *UserFindByIdParser) SetParser(interface{}) {
	return
}

func (p *UserFindByIdParser) GetModel() *model.User {
	return &p.model
}

func (p *UserFindByIdParser) TryParse(data string) error {
	(*UserIdFormatParser).TryParse(nil, 
}

func NewOutputParser() *gateway.Output {
	return &gateway.Output{}
}

/*
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
*/
