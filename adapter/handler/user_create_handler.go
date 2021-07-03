/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_create_handler.go
UserCreateHandler
*/

package handler

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"anemone/adapter/helper"
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

type UserCreateHandler interface {
	Handler
	Handle() *cobra.Command
	SetUseCase(interface{})
	Create() error
}

type userCreateHandler struct {
	interactor usecase.IUserUseCase
	rhandler
}

//func NewUserCreateHandler(u usecase.IUserUseCase) UserCreateHandler {
func NewUserCreateHandler() UserCreateHandler {
	r := &userCreateHandler{}
	r.AddSetter("UserUseCase", r.SetUserUseCase)
	r.SetHandle(r.Handle)
	return r
}

func init() {
	_, err := helper.GetBuilderInstance("user")
	//construct := Constructor("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	//handler := NewUserCreateHandler()
	//construct.register(handler.GetHandle())
}

func (s *userCreateHandler) SetUseCase(u interface{}) {
	s.interactor = u.(usecase.IUserUseCase)
}

func (s *userCreateHandler) Handle() *cobra.Command {

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
			_, _ = cmd.Flags().GetString("data")
			err := s.Create()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "Your name")
	return cmd
}

func (s *userCreateHandler) Create() error {

	err := s.interactor.Create()
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}
