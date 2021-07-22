// +build user

/* vim: set ts=4 sw=4: */
/*
adapter/handler/user_remove_handler.go
UserRemoveHandler
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
// 2021/06/27 UserRemoveHandler
// 引数が指定されたフォーマットじゃない場合はエラーのバリデーションが未実装
//
// 2021/06/27 UserRemoveHandler
// 引数のidに指定された文字列以外が使われてる場合はエラーのバリデーションが未実装
//
// HISTORY(koube):
// 2021/06/27 UserRemoveHandler
// 新規作成
//
// 2021/07/05 UserRemoveHandler
// UseCaseSetterとCommandConstructorに対応
//
// 2021/07/26 UserRemoveHandler
// 指定された引数のバリデーションをParserクラスに切り出す。
// Parserはインジェクションで実装を入れ替えられる。外界との変換処理は
// Parserクラスの担当として実装する。

// UserRemoveHanlder のinterface定義
type UserRemoveHanlder interface {
	Handler
	SetUseCase(interface{})
	Remove(id string) error
}

type userRemoveHandler struct {
	interactor usecase.IUserUseCase
	*UserRemoveParser
	rhandler
}

func NewUserRemoveHandler(p *UserRemoveParser) UserRemoveHanlder {
	r := &userRemoveHandler{UserRemoveParser: p}
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
	format := gateway.NewUserIdFormatParser()

	parser := NewUserRemoveParser(format)
	handler := NewUserRemoveHandler(parser)
	constructor.Register(handler.GetHandle())
}

func (s *userRemoveHandler) SetUseCase(u interface{}) {
	s.interactor = u.(usecase.IUserUseCase)
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
			// idが指定されたフォーマットじゃない場合はエラー
			parser, ok := interface{}(s).(gateway.IFormatParser)
			if ok {
				err := parser.TryParse(id, s.GetModel())
				if err != nil {
					return err
				}
				return nil
			}
			// TryParserが存在しない（インターフェースを満たさない）場合は
			// Internal Server Error
			return errors.New(
				codes.InternalServerError,
				"Required Method dose not exist in the specified parser",
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetString("id")
			if err := s.Remove(id); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("id", "", "Specify the ID that identifies the User")
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
	s.gateway.SetResponse(
		"",
		/*
			レスポンスはなし。正常に終了した事を表す何か？
			nilをResponseにセットする？
		*/)
	return nil
}
