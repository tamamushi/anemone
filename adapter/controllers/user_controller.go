/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	_ "anemone/adapter/handler"
	"anemone/adapter/helper"
	"anemone/codes"
	"anemone/errors"
)

// TODO(koube):
// 2021/06/27 UserController
// GetBuilderInstanceに絡む部分のテストが未実施。正常なBuilderインスタンスが
// 返って来なかった場合はエラーが反る。
//
// HISTORY(koube):
// 2021/06/27 UserController 新規作成

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

type userController struct {
}

// UserCoontroller のコンストラクタ
func NewUserController() IUserController {
	return &userController{}
}

var commandBuilder = helper.GetBuilderInstance

func init() {
	blder, err := commandBuilder("root")
	if err != nil {
		msg := "Failed to building Root command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	controller := NewUserController()
	blder.AddCommand(controller.Handler())
}

func (s *userController) Handler() *cobra.Command {

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

	// 子ハンドラ達を取り出す。
	// usecaseにDBを設定し、子ハンドラにわたす。
	// DBインスタンスとusecaseを持った子ハンドラの
	// cobraインスタンスをAddCommandで登録する

	blder, err := commandBuilder("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}

	// for _, handler := range blder.GetCommands() {
	//		s.infra
	//		handler.SetUseCase(usecase)
	// 		cmd.AddCommand(handler.Habdle())
	//}

	cmd.AddCommand(blder.GetCommands()...)
	return cmd
}
