/* vim: set ts=4 sw=4: */

package test

import (
	"anemone/adapter/gateway"
)

type parserMock struct {
	gateway.IParser
	MockTryParseFormat func(string) error
	MockInput          func(string) interface{}
}

func NewParserMock(m interface{}) *parserMock {
	return &parserMock{}
}

func (p *parserMock) TryParseFormat(s string) error {
	return p.MockTryParseFormat(s)
}
func (p *parserMock) Input(s string) interface{} {
	return p.MockInput(s)
}

type ParserMethod struct {
	TryParseFormat func(string) error
	Input          func(string) interface{}
}

func GetParserMethodStruct() *ParserMethod {
	return &ParserMethod{nil, nil}
}
func (p *ParserMethod) SetTryParseFormat(
	f func(string) error) *ParserMethod {
	p.TryParseFormat = f
	return p
}
func (p *ParserMethod) SetInput(
	f func(string) interface{}) *ParserMethod {
	p.Input = f
	return p
}
