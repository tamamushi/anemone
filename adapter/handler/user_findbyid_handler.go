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
// 引数のidに指定された文字列以外が使われてる場合はエラーの
// バリデーションが未実装
//
// 2021/07/25 UserFindByIdHandler
// 指定された引数のバリデーションはParserクラスに切り出す。
// Parserはインジェクションで実装を入れ替える。Output処理もParserに統合し
// 外界との変換処理はParserクラスの担当として実装する。
//
// HISTORY(koube):
// 2021/06/27 UserFindByIdHandler 新規作成
// 2021/07/05 UserFindByIdHandler UseCaseSetterとCommandConstructorに対応
// 2021/07/26 UserFindByIdHandler UserFindByIdParserに対応させる

// UserFindByIdHandler のinterface定義
type UserFindByIdHandler interface {
	Handler
	SetUseCase(interface{})
	FindById(id string) error
}

type userFindByIdHandler struct {
	interactor usecase.IUserUseCase
	parser     *UserFindByIdParser
	rhandler
}

//func NewUserFindByIdHandler(p gateway.IParser) UserFindByIdHandler {
func NewUserFindByIdHandler(p *UserFindByIdParser) UserFindByIdHandler {
	s := &userFindByIdHandler{parser: p}
	s.AddSetter("UserUseCase", s.SetUseCase)
	s.SetHandle(s.Handle)
	return s
}

func init() {
	constructor, err := Constructor("user")
	FatalConstruction(
		err,
		"Failed to building User command group (%s)",
	)
	format := gateway.NewUserIdFormatParser()
	output := gateway.NewUserOutputParser()

	parser := NewUserFindByIdParser(format, output)
	handler := NewUserFindByIdHandler(parser)
	constructor.Register(handler.GetHandle())
}

func (s *userFindByIdHandler) SetUseCase(u interface{}) {
	s.interactor = u.(usecase.IUserUseCase)
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
			// idが指定されたフォーマットじゃない場合はエラー
			parser, ok := interface{}(s.parser).(gateway.IFormatParser)
			if ok {
				err := parser.TryParse(id, s.parser.GetModel())
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
			// 見つかったユーザー情報の返却
			if err := s.FindById(id); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("id", "", "Specify the ID that identifies the User")
	return cmd
}

func (s *userFindByIdHandler) FindById(id string) error {
	result, err := s.interactor.FindById(id)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	output, ok := interface{}(s.parser).(gateway.IOutput)
	if !ok {
		// Outputが存在しない（インターフェースを満たさない）場合は
		// Internal Server Error
		return errors.New(
			codes.InternalServerError,
			"Required Method dose not exist in the specified parser",
		)
	}
	s.gateway.SetResponse(output.Output(result))
	return nil
}
