/* vim: set ts=4 sw=4: */

package controllers_test

import (
	"anemone/adapter/controllers"
	"anemone/adapter/gateway"
	"anemone/application/usecase"
)

func Example() {
	// 初期化サンプル
	// init()の伝播の中でコマンド構築を行う

	blder, err := controllers.CommandBuilder("root")
	controllers.FatalBuilder(
		err,
		"Failed to building Root command group (%s)",
	)
	interactor := usecase.NewUserInteractor()
	controller := controllers.NewUserController(interactor)
	gateway := gateway.GetGateway()
	blder.AddCommand(controller.Handler(gateway))
}
