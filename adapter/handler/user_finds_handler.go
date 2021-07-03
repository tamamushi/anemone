/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_finds_handler.go
UserFindsHandler
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
// 2021/06/27 UserFindsHandler
// 検索パラメータの指定方法が確定していない。
// パラメータ指定方法が確定していないのでバリデーションは全て未実装。
// 考えられる候補は limit、keyword、date、flags/tags
//
// HISTORY(koube):
// 2021/06/27 UserFindsHandler 新規作成

// UserFindsHandler のinterface定義
type UserFindsHandler interface {
	Handle() *cobra.Command
	Finds() error
}

type userFindsHandler struct {
	interactor usecase.IUserUseCase
}

func NewUserFindsHandler(u usecase.IUserUseCase) UserFindsHandler {
	return &userFindsHandler{u}
}

func init() {
	blder, err := helper.GetBuilderInstance("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	handler := NewUserFindsHandler(usecase)
	blder.AddCommand(handler.Handle())
}

func (s *userFindsHandler) Handle() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "finds",
		Short: "User Finds command",
		Args: func(cmd *cobra.Command, args []string) error {
			/* パラメータの候補 limit、 keyword、 date、 flags/tags */
			/*
				// id が指定されていなければエラー
				if len(id) < 1 {
					return errors.New(
						codes.NotEnoughArgument,
						"Need to specified ID",
					)
				}
			*/
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
			//id, _ := cmd.Flags().GetString("id")
			// 見つかったユーザー情報の返却
			err := s.Finds()
			if err != nil {
				return err
			}
			return nil
		},
	}
	//cmd.Flags().String("id", "", "Your name")
	return cmd
}

func (s *userFindsHandler) Finds() error {

	err := s.interactor.Finds()
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}
