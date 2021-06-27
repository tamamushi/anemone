/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_remove_handler.go
UserRemoveHandler
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
// 2021/06/27 UserRemoveHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserRemoveHandler
// 引数のidに指定された文字列以外が使われてる場合はエラーのバリデーションが未実装
//
// HISTORY(koube):
// 2021/06/27 UserRemoveHandler 新規作成

// UserRemoveHanlder のinterface定義
type UserRemoveHanlder interface {
	Handle() *cobra.Command
	Remove(id string) error
}

type userRemoveHandler struct {
	interactor usecase.IUserUseCase
}

func NewUserRemoveHandler(u usecase.IUserUseCase) UserRemoveHanlder {
	return &userRemoveHandler{u}
}

func init() {
	blder, err := helper.GetBuilderInstance("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	handler := NewUserRemoveHandler(usecase)
	blder.AddCommand(handler.Handle())
}

func (s *userRemoveHandler) Handle() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "User Remove Command",
		Args: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetString("id")
			// id が指定されていなければエラー
			if len(id) < 1 {
				return errors.New(
					codes.NotEnoughArgument,
					"Need to specified ID",
				)
			}
			// id の桁数が指定されたフォーマットじゃない場合はエラー
			/*
				if len(id) > 20 {
					return errors.New(
						codes.InvalidArgument,
						"Allow the id formats XXXX-XXXX-XXXX-XXXX",
					)
				}
			*/
			// id が指定されたキャラクタセットじゃなければエラー
			// キャラクタセットは、0-9、A-Z（小文字のa-zは含まない）
			/*
				if len(id) > 20 {
					return errors.New(
						codes.InvalidArgument,
						"Allow the usable character is 0-9, A-Z",
					)
				}
			*/
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetString("id")
			err := s.Remove(id)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("id", "", "Your name")
	return cmd
}

func (s *userRemoveHandler) Remove(id string) error {

	err := s.interactor.Remove(id)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}
