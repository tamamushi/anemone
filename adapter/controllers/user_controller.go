/* vim: set ts=4 sw=4: */

/*
Controllers は外界からのインプットの内側へのルーティング、
内側からの外界へのレスポンスを担います。
CleanArchitectureのInterface Adaptersにおける振る舞いを
実装しています。
*/
package controllers

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	_ "anemone/adapter/handler"
	"anemone/adapter/helper"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// TODO(koube):
// GetBuilderInstanceに絡む部分のテストが未実施。正常なBuilderインスタンスが
// 返って来なかった場合はエラーが反る。

// Userコマンドを構築する起点になるController
// 実行可能なコマンドは以下がある
//
//  create
//  delete
//  update
//  findbyid
//  finds
//
// Interface Definition
//
// UserControllerのInterfaceは定義は以下の通り
type IUserController interface {
	Handler() *cobra.Command
}

type controller struct {
	interactor usecase.IUserUseCase
}

// UserCoontroller のコンストラクタ
func NewUserController(u usecase.IUserUseCase) IUserController {
	return &controller{u}
}

var commandBuilder = helper.GetBuilderInstance

func init() {
	blder, err := commandBuilder("root")
	if err != nil {
		msg := "Failed to building Root command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	controller := NewUserController(usecase)
	blder.AddCommand(controller.Handler())
}

func (s *controller) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "User handle command group",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				// Error コードを返す
				return errors.New(
					codes.NotEnoughArgument,
					"Required target sub command",
				)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New(
				codes.UnSupportedMethod,
				errors.Messagef("UnSupported called method: %s", args[0]),
			)
		},
	}

	blder, err := commandBuilder("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	cmd.AddCommand(blder.GetCommands()...)
	return cmd
}
