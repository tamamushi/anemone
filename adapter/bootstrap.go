/* vim: set ts=4 sw=4: */

/*
Adapter は実行時のブートストラップが含まれる。

CleanArchitectureのInterface Adaptersにおける振る舞いを実装しています。
*/
package adapter

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"

	"anemone/adapter/controllers"
	"anemone/adapter/gateway"
	"anemone/adapter/presenter"
)

type Response struct {
	Err error
	Cmd *cobra.Command
}

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "anemone",
		Short: "A brief description of your application",
	}
	cobra.OnInitialize(initConfig)

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return cmd
}

func Bootstrap(p presenter.Presenter) {

	cmd := NewCmdRoot()
	buffer := &bytes.Buffer{}
	cmd.SetOut(buffer)
	cmd.SetErr(buffer)

	// DBインスタンスをSingletonで登録
	// Controllerとして分けられた各コマンドは
	// 登録されたDBインスタンスのポインタを使って
	// 自分の子ハンドラに設定する。

	// Controllerは取り込んだDBインスタンスを
	// 各子ハンドラへusecase設定と共に設定する

	blder, err := controllers.CommandBuilder("root")
	if ok := controllers.FatalBuilder(
		err,
		"Failed to building Root command group (%s)",
	); ok {
		// CommandBuilderはanemone空間のコマンドを保持する
		// blder.GetCommandsでanemone空間のコマンドを展開し
		// cobraのコマンド実行プロセスに登録する
		cmd.AddCommand(blder.GetCommands()...)
		cmd.SetArgs(append(p.Command(), p.Arguments()...))
		// コマンド実行はcobraのコマンド実行プロセスを流用する
		cobra.CheckErr(cmd.Execute())

		// gatewayはシングルトンでcontrollers内でインスタンス化され、
		// 実行結果の出力情報などを保持する。
		gateway := gateway.GetGateway()
		cmd.Printf("%s", gateway.Response())
		fmt.Print(buffer)
		// gatewayを返す仕様にする？
	}
}

func initConfig() {
}
