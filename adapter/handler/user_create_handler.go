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
// 2021/07/17 UserCreateHandler
// 引数の引渡し方式確定。バリデーションも実装。
// バリデーション処理自体はParserクラスに切り出し
// Create処理成功時のレスポンス処理が未実装。
//
// HISTORY(koube):
// 2021/06/27 UserCreateHandler 新規作成
// 2021/07/05 UserCreateHandler UseCaseSetterとCommandConstructorに対応
// 2021/07/17 UserCreateHandler UserCreateParserに対応

type UserCreateHandler interface {
	Handler
	SetUseCase(interface{})
	Create(string) error
}

type userCreateHandler struct {
	interactor usecase.IUserUseCase
	parser     gateway.IParser
	rhandler
}

func NewUserCreateHandler(p gateway.IParser) UserCreateHandler {
	s := &userCreateHandler{parser: p}
	s.AddSetter("UserUseCase", s.SetUseCase)
	s.SetHandle(s.Handle)
	return s
}

func init() {
	construct, err := Constructor("user")
	FatalConstruction(
		err,
		"Failed to building User command group (%s)",
	)
	// CreateHandlerに対応する。つまりUserUseCaseの
	// Createに対するInputPortを動的に提供する先を
	// 生成する。
	parser := NewUserCreateParser()
	handler := NewUserCreateHandler(parser)
	construct.Register(handler.GetHandle())
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
			// data が指定されていなければエラー
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
			if err := s.Create(data); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "Your name")
	return cmd
}

func (s *userCreateHandler) Create(data string) error {
	user, err := s.interactor.Create(s.parser.Input(data))
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
