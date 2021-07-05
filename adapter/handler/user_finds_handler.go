// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_finds_handler.go
UserFindsHandler
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
// 2021/06/27 UserFindsHandler
// 検索パラメータの指定方法が確定していない。
// パラメータ指定方法が確定していないのでバリデーションは全て未実装。
// 考えられる候補は limit、keyword、date、flags/tags
//
// HISTORY(koube):
// 2021/06/27 UserFindsHandler 新規作成
// 2021/07/05 UserFindsHandler UseCaseSetterとCommandConstructorに対応

// UserFindsHandler のinterface定義
type UserFindsHandler interface {
	Handler
	SetUseCase(interface{})
	Finds() error
}

type userFindsHandler struct {
	interactor usecase.IUserUseCase
	rhandler
}

func NewUserFindsHandler() UserFindsHandler {
	r := &userFindsHandler{}
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
	handler := NewUserFindsHandler()
	constructor.Register(handler.GetHandle())
}

func (s *userFindsHandler) SetUseCase(u interface{}) {
	s.interactor = u.(usecase.IUserUseCase)
}

func (s *userFindsHandler) Handle(g gateway.Gateway) *cobra.Command {

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
