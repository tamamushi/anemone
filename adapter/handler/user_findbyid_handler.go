// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_findbyid_handler.go
UserFindByIdHandler
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
// 2021/06/27 UserFindByIdHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserFindByIdHandler
// 引数のidに指定された文字列以外が使われてる場合はエラーのバリデーションが未実装
//
// HISTORY(koube):
// 2021/06/27 UserFindByIdHandler 新規作成
// 2021/07/05 UserFindByIdHandler UseCaseSetterとCommandConstructorに対応

// UserFindByIdHandler のinterface定義
type UserFindByIdHandler interface {
	Handler
	SetUseCase(interface{})
	FindById(id string) error
}

type userFindByIdHandler struct {
	interactor usecase.IUserUseCase
	rhandler
}

func NewUserFindByIdHandler() UserFindByIdHandler {
	r := &userFindByIdHandler{}
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
	handler := NewUserFindByIdHandler()
	constructor.Register(handler.GetHandle())
}

func (s *userFindByIdHandler) SetUseCase(u interface{}) {
	s.interactor = u.(usecase.IUserUseCase)
}

func (s *userFindByIdHandler) Handle(g gateway.Gateway) *cobra.Command {

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
