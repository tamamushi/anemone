// +build user

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
// 2021/07/26 UserUpdateHandler
// OutputPortによる変換処理失敗時のエラー処理が未実装
//
// HISTORY(koube):
// 2021/06/27 UserUpdateHandler
// 新規作成
//
// 2021/07/05 UserUpdateHandler
// UseCaseSetterとCommandConstructorに対応
//
// 2021/07/17 UserUpdateHandler
// 引数の引渡し方式確定。バリデーションも実装。
// バリデーション処理自体はParserクラスに切り出し
// Update処理成功時のレスポンス処理が未実装。
//
// 2021/07/26 UserUpdateHandler UserUpdateParserに対応。指定された引数の
// バリデーションをParserクラスに切り出す。Parserはインジェクションで実装を
// 入れ替える。Output処理もParserに統合し外界との変換処理はParserクラスの担当
// として実装する。

type UserUpdateHandler interface {
	Handler
	SetUseCase(interface{})
	Update(string) error
}

type userUpdateHandler struct {
	interactor usecase.IUserUseCase
	*UserUpdateParser
	rhandler
}

func NewUserUpdateHandler(p *UserUpdateParser) UserUpdateHandler {
	r := &userUpdateHandler{UserUpdateParser: p}
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
	format := gateway.NewUserModelFormatParser()
	input := gateway.NewUserInputParser()
	output := gateway.NewUserOutputParser()

	parser := NewUserUpdateParser(format, input, output)
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
			parser, ok := interface{}(s).(gateway.IFormatParser)
			if ok {
				err := parser.TryParse(data, s.GetModel())
				if err != nil {
					return errors.New(
						codes.InvalidArgument,
						errors.Messagef(
							"Argument faild parse. DATA format invalid %v#",
							data,
						),
					)
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
			data, _ := cmd.Flags().GetString("data")
			if err := s.Update(data); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "Specify the data in JSON format")
	return cmd
}

func (s *userUpdateHandler) Update(data string) error {

	inputPort, ok := interface{}(s).(gateway.IInput)
	if !ok {
		// Inputが存在しない（インターフェースを満たさない）場合は
		// Internal Server Error
		return errors.New(
			codes.InternalServerError,
			"Required Method:[Input()] dose not exist in the specified parser",
		)
	}
	result, err := s.interactor.Update(
		inputPort.Input(data, s.GetModel()),
	)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}

	outputPort, ok := interface{}(s).(gateway.IOutput)
	if !ok {
		// Outputが存在しない（インターフェースを満たさない）場合は
		// Internal Server Error
		return errors.New(
			codes.InternalServerError,
			"Required Method:[Output] dose not exist in the specified parser",
		)
	}
	s.gateway.SetResponse(outputPort.Output(result))
	return nil
}
