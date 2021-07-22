/* vim: set ts=4 sw=4: */

/*
TryParse、Input、Outputを持ち、Handlerに必要なメソッドの実装を提供する。
HandlerによってTryParseやInputの処理は可変となり、コンストラクタで実装を
設定してインスタンスを返す。
*/
package gateway

import (
	"anemone/codes"
	"anemone/errors"
	"encoding/json"
)

// UserParserはUser Modelを外界との受渡し処理に関する実装。
// IParserインターフェースの実装
type UserParser struct {
	funcTryParse func(string, interface{}) error
	funcInput    func(string, interface{}) interface{}
	funcOutput   func(interface{}) string
}

func NewUserModelFormatParser() IFormatParser {
	p := &UserParser{}
	p.funcTryParse = tryParseData
	return p
}

func NewUserIdFormatParser() IFormatParser {
	p := &UserParser{}
	p.funcTryParse = tryParseId
	return p
}

func NewUserInputParser() IInput {
	p := &UserParser{}
	p.funcInput = inputModel
	return p
}

func NewUserOutputParser() IOutput {
	p := &UserParser{}
	p.funcOutput = outputModel
	return p
}

func (p *UserParser) TryParse(d string, m interface{}) error {
	return p.funcTryParse(d, m)
}

func (p *UserParser) Output(m interface{}) string {
	return p.funcOutput(m)
}

func (p *UserParser) Input(d string, m interface{}) interface{} {
	return p.funcInput(d, m)
}

// 以下は実体
// modelが返された場合のOutputPort。Stringへ変換
func outputModel(m interface{}) string {
	model, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(model)
}

// stringでJSONが渡された場合のInputPort。Modelへ変換
func inputModel(data string, m interface{}) interface{} {
	if err := json.Unmarshal([]byte(data), m); err != nil {
		return err
	}
	return m
}

// dataオプション用のTryParse
func tryParseData(data string, m interface{}) error {
	if err := json.Unmarshal([]byte(data), m); err != nil {
		return err
	}
	return nil
}

// Idオプション用のTryParse
func tryParseId(id string, _ interface{}) error {
	// id の桁数が指定されたフォーマットじゃない場合はエラー
	if len(id) > 20 {
		return errors.New(
			codes.InvalidArgument,
			"Allow the id formats XXXX-XXXX-XXXX-XXXX",
		)
	}
	// id が指定されたキャラクタセットじゃなければエラー
	// キャラクタセットは、0-9、A-Z（小文字のa-zは含まない）
	if len(id) > 20 {
		return errors.New(
			codes.InvalidArgument,
			"Allow the usable character is 0-9, A-Z",
		)
	}
	return nil
}
