/* vim: set ts=4 sw=4: */

package handler_test

import (
	"fmt"
	"os"

	"anemone/adapter/handler"
	"anemone/adapter/helper"
	"anemone/application/usecase"
)

func Example() {
	// 初期化サンプル
	// init()の伝播の中でコマンド構築を行う

	blder, err := helper.GetBuilderInstance("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	handler := handler.NewUserCreateHandler(usecase)
	blder.AddCommand(handler.Handle())
}

// テスト番号出力のためのグローバル変数
var index = 0
