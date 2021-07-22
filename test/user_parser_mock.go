/* vim: set ts=4 sw=4: */

package test

import (
	"anemone/adapter/gateway"
	"anemone/model"
)

type ParserMock struct {
	gateway.IFormatParser
	gateway.IOutput
	gateway.IInput
	MockTryParse func(string, interface{}) error
	MockInput    func(string, interface{}) interface{}
	MockOutput   func(interface{}) string
}

func NewParserMock() *ParserMock {
	return &ParserMock{}
}

func (p *ParserMock) TryParse(s string, m interface{}) error {
	return p.MockTryParse(s, m)
}

func (p *ParserMock) Output(m interface{}) string {
	return p.MockOutput(m)
}

func (p *ParserMock) Input(s string, m interface{}) interface{} {
	return p.MockInput(s, m)
}

type ParserMethod struct {
	TryParse func(string, interface{}) error
	Input    func(string, interface{}) interface{}
	Output   func(interface{}) string
}

func GetParserMethodStruct() *ParserMethod {
	return &ParserMethod{nil, nil, nil}
}

func (p *ParserMethod) SetTryParse(
	f func(string, interface{}) error) *ParserMethod {
	p.TryParse = f
	return p
}

func (p *ParserMethod) SetInput(
	f func(string, interface{}) interface{}) *ParserMethod {
	p.Input = f
	return p
}

func (p *ParserMethod) SetOutput(
	f func(interface{}) string) *ParserMethod {
	p.Output = f
	return p
}

func PrepareParserMock(tt *TCase) gateway.IFormatParser {
	format := NewParserMock()
	method, _ := tt.GetParser()
	inter, ok := method.(*ParserMethod)
	if ok && inter.TryParse != nil {
		format.MockTryParse = inter.TryParse
	} else {
		format.MockTryParse = func(s string, _ interface{}) error { return nil }
	}
	return format
}

func PrepareInputMock(tt *TCase) gateway.IInput {
	input := NewParserMock()
	method, _ := tt.GetParser()
	inter, ok := method.(*ParserMethod)
	if ok && inter.Input != nil {
		input.MockInput = inter.Input
	} else {
		input.MockInput = func(d string, m interface{}) interface{} {
			return &model.User{}
		}
	}
	return input
}

func PrepareOutputMock(tt *TCase) gateway.IOutput {
	output := NewParserMock()
	method, _ := tt.GetParser()
	inter, ok := method.(*ParserMethod)
	if ok && inter.Output != nil {
		output.MockOutput = inter.Output
	} else {
		output.MockOutput = func(m interface{}) string { return "dummy" }
	}
	return output
}
