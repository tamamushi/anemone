/* vim: set ts=4 sw=4: */

/*
Adapter は実行時のブートストラップが含まれる。
CleanArchitectureのInterface Adaptersにおける振る舞いを
実装しています。
*/
package adapter

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"

	"anemone/adapter/controllers"
	"anemone/adapter/helper"
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

func Bootstrap() {

	cmd := NewCmdRoot()
	buffer := &bytes.Buffer{}
	cmd.SetOutput(buffer)

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
		cmd.AddCommand(blder.GetCommands()...)

		args := helper.NewArgumentBuilder()
		args.AddCommand("user", "create")

		jsond := "{ \"id\": \"xxxx-xxxx-xxxx-xxx\", \"email\": \"t.koube.cp@gmail.com\" }"
		cmd.SetArgs([]string{"user", "update", "--data", jsond})
		//fmt.Printf("%s", []string{"user", "create", "--data", jsond})

		cobra.CheckErr(cmd.Execute())

		fmt.Print(buffer)
	}
}

func initConfig() {
}
