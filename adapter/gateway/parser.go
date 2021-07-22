/* vim: set ts=4 sw=4: */

/*
Parser

ParserはHandlerに外界との処理実装を提供する。
引数のバリデーション、UseCaseとのmodelの受渡しに必要な変換処理を提供する、
Handlerのサブモジュール。
*/
package gateway

type IParser interface {
	SetParser(interface{})
	GetModel() interface{}
}

type IFormatParser interface {
	TryParse(string, interface{}) error
}

type IInput interface {
	Input(string, interface{}) interface{}
}

type IOutput interface {
	Output(interface{}) string
}
