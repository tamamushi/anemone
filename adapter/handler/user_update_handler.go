/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_update_handler.go
UserUpdateHandler
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
// 2021/06/27 UserUpdateHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserUpdateHandler
// 引数をInteractorにデータとして渡す方式が固まってない。その為引渡処理が未実装
//
//
// HISTORY(koube):
// 2021/06/27 UserUpdateHandler 新規作成

type UserUpdateHandler interface {
	Handle() *cobra.Command
	Update() error
}

type userUpdateHandler struct {
	interactor usecase.IUserUseCase
}

func NewUserUpdateHandler(u usecase.IUserUseCase) UserUpdateHandler {
	return &userUpdateHandler{u}
}

func init() {
	blder, err := helper.GetBuilderInstance("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	handler := NewUserUpdateHandler(usecase)
	blder.AddCommand(handler.Handle())
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
