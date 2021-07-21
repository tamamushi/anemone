/* vim: set ts=4 sw=4: */

package handler_test

import (
	_ "anemone/adapter/handler"
)

func Example() {
	// 初期化サンプル
	// init()の伝播の中でコマンド構築を行う

	/*
		constructor, err := handler.Constructor("user")
		handler.FatalConstruction(
			err,
			"Failed to building User command group (%s)",
		)
			parser := handler.NewUserCreateParser()
			handler := handler.NewUserCreateHandler(parser)
		constructor.Register(handler.GetHandle())
	*/
}

// テスト番号出力のためのグローバル変数
var index = 0
