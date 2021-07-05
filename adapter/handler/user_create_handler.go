// +build user

/* vim: set ts=4 sw=4:
adapter/handler/user_create_handler.go
UserCreateHandler
*/

package handler

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// TODO(koube):
// 2021/06/27 UserCreateHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserCreateHandler
// 引数をInteractorにデータとして渡す方式が固まってない。その為引渡処理が未実装
//
// HISTORY(koube):
// 2021/06/27 UserCreateHandler 新規作成
// 2021/07/05 UserCreateHandler UseCaseSetterとCommandConstructorに対応

type UserCreateHandler interface {
	Handler
	SetUseCase(interface{})
	Create(string) error
}

type userCreateHandler struct {
	interactor usecase.IUserUseCase
	rhandler
}

func NewUserCreateHandler() UserCreateHandler {
	r := &userCreateHandler{}
	r.AddSetter("UserUseCase", r.SetUseCase)
	r.SetHandle(r.Handle)
	return r
}

func init() {
	construct, err := Constructor("user")
	FatalConstruction(
		err,
		"Failed to building User command group (%s)",
	)
	handler := NewUserCreateHandler()
	construct.Register(handler.GetHandle())
}

func (s *userCreateHandler) SetUseCase(u interface{}) {
	s.interactor = u.(usecase.IUserUseCase)
}

func (s *userCreateHandler) Handle(g gateway.Gateway) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "User Create Command",
		Args: func(cmd *cobra.Command, args []string) error {
			data, _ := cmd.Flags().GetString("data")
			// id が指定されていなければエラー
			if len(data) < 1 {
				return errors.New(
					codes.NotEnoughArgument,
					"Need to specified DATA",
				)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := cmd.Flags().GetString("data")
			if err != nil {
				return err
			}
			err = s.Create(data)
			cmd.Printf("create success!")
			cmd.Printf("%s", g.Hoge())
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "Your name")
	return cmd
}

func (s *userCreateHandler) Create(_ string) error {

	err := s.interactor.Create(
	//s.prenseter.Input(c)
	//s.gateway.Output()
	)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}
