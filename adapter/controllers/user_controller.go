// +build user

/* vim: set ts=4 sw=4:
adapter/controllers/user_controller.go
UserController
*/

package controllers

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	"anemone/adapter/handler"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// TODO(koube):
// 2021/06/27 UserController
// CommandBuilderに絡む部分のテストが未実施。正常なBuilderインスタンスが
// 返って来なかった場合はエラーが反る。
//
// HISTORY(koube):
// 2021/06/27 UserController 新規作成
// 2021/07/05 UserController UseCaseSetterに対応

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
	Handler(gateway.IGateway) *cobra.Command
}

type userController struct {
	interactor usecase.IUserUseCase
}

// UserCoontroller のコンストラクタ
func NewUserController(u usecase.IUserUseCase) IUserController {
	return &userController{u}
}

func init() {
	blder, err := CommandBuilder("root")
	FatalBuilder(
		err,
		"Failed to building Root command group (%s)",
	)
	interactor := usecase.NewUserInteractor()
	controller := NewUserController(interactor)
	gateway := gateway.GetGateway()
	blder.AddCommand(controller.Handler(gateway))
}

func (s *userController) Handler(g gateway.IGateway) *cobra.Command {

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
	handlers, err := handler.Constructor("user")
	if ok := handler.FatalConstruction(
		err,
		"Failed to building User command group (%s)",
	); ok {
		for _, handle := range handlers.Extraction() {
			for k, v := range handle.GetSetters() {
				switch k {
				case "UserUseCase":
					v(s.interactor)
				}
			}
			handle.SetGateway(g)
			cmd.AddCommand(handle.Handle())
		}
	}
	return cmd
}
