/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_findbyid_handler.go
UserFindByIdHandler
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
// 2021/06/27 UserFindByIdHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserFindByIdHandler
// 引数のidに指定された文字列以外が使われてる場合はエラーのバリデーションが未実装
//
// HISTORY(koube):
// 2021/06/27 UserFindByIdHandler 新規作成

// UserFindByIdHandler のinterface定義
type UserFindByIdHandler interface {
	Handle() *cobra.Command
	FindById(id string) error
}

type userFindByIdHandler struct {
	interactor usecase.IUserUseCase
}

func NewUserFindByIdHandler(u usecase.IUserUseCase) UserFindByIdHandler {
	return &userFindByIdHandler{u}
}

func init() {
	blder, err := helper.GetBuilderInstance("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	handler := NewUserFindByIdHandler(usecase)
	blder.AddCommand(handler.Handle())
}

func (s *userFindByIdHandler) Handle() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "findbyid",
		Short: "User FindById command",
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
			// 見つかったユーザー情報の返却
			err := s.FindById(id)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("id", "", "Your name")
	return cmd
}

func (s *userFindByIdHandler) FindById(id string) error {

	err := s.interactor.FindById(id)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}
