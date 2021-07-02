/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_update_handler.go
UserUpdateHandler
*/

package handler

import (
	"github.com/spf13/cobra"

	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// TODO(koube):
// 2021/06/27 UserUpdateHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserUpdateHandler
// 引数をInteractorにデータとして渡す方式が固まってない。その為引渡処理が未実装
//
// HISTORY(koube):
// 2021/06/27 UserUpdateHandler 新規作成
// 2021/07/05 UserUpdateHandler UseCaseSetterとCommandConstructorに対応

type UserUpdateHandler interface {
	Handler
	SetUseCase(interface{})
	Update() error
}

type userUpdateHandler struct {
	interactor usecase.IUserUseCase
	rhandler
}

func NewUserUpdateHandler() UserUpdateHandler {
	r := &userUpdateHandler{}
	r.AddSetter("UserUseCase", r.SetUseCase)
	r.SetHandle(r.Handle)
	return r
}

func init() {
	constructor, err := Constructor("user")
	FatalConstruction(
		err,
		"Failed to building User command group (%s)",
	)
	handler := NewUserUpdateHandler()
	constructor.Register(handler.GetHandle())
}

func (s *userUpdateHandler) SetUseCase(u interface{}) {
	s.interactor = u.(usecase.IUserUseCase)
}

func (s *userUpdateHandler) Handle() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "update",
		Short: "User Update Command",
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
			_, _ = cmd.Flags().GetString("data")
			err := s.Update()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "Your name")
	return cmd
}

func (s *userUpdateHandler) Update() error {

	err := s.interactor.Update()
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}
