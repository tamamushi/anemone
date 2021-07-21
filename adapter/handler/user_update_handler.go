// +build -user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_update_handler.go
UserUpdateHandler
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
// 2021/06/27 UserUpdateHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserUpdateHandler
// 引数をInteractorにデータとして渡す方式が固まってない。その為引渡処理が未実装
//
// 2021/07/17 UserUpdateHandler
// 引数の引渡し方式確定。バリデーションも実装。
// バリデーション処理自体はParserクラスに切り出し
// Update処理成功時のレスポンス処理が未実装。
//
// HISTORY(koube):
// 2021/06/27 UserUpdateHandler 新規作成
// 2021/07/05 UserUpdateHandler UseCaseSetterとCommandConstructorに対応
// 2021/07/17 UserUpdateHandler UserUpdateParserに対応

type UserUpdateHandler interface {
	Handler
	SetUseCase(interface{})
	Update(string) error
}

type userUpdateHandler struct {
	interactor usecase.IUserUseCase
	parser     gateway.IParser
	rhandler
}

func NewUserUpdateHandler(p gateway.IParser) UserUpdateHandler {
	r := &userUpdateHandler{parser: p}
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
	parser := NewUserUpdateParser()
	handler := NewUserUpdateHandler(parser)
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
			// data が所定のフォーマット（JSON形式）でなければエラー
			if err := s.parser.TryParseFormat(data); err != nil {
				return errors.New(
					codes.InvalidArgument,
					errors.Messagef(
						"Argument faild parse. DATA format invalid %v#",
						data,
					),
				)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			data, _ := cmd.Flags().GetString("data")
			if err := s.Update(data); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "target data by json format")
	return cmd
}

func (s *userUpdateHandler) Update(data string) error {

	user, err := s.interactor.Update(s.parser.Input(data))
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	s.gateway.SetResponse(user)
	return nil
}
