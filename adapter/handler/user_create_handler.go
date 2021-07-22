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
// Create処理成功時のレスポンス処理が未実装
//
// 2021/07/26 UserCreateHandler
// OutputPortによる変換処理失敗時のエラー処理が未実装
//
// HISTORY(koube):
// 2021/06/27 UserCreateHandler
// 新規作成
//
// 2021/07/05 UserCreateHandler
// UseCaseSetterとCommandConstructorに対応
//
// 2021/07/17 UserCreateHandler UserCreateParserに対応。引数の引渡し方式確定。
// バリデーションも実装。バリデーション処理自体はParserクラスに切り出し
//
// 2021/07/26 UserCreateHandler UserCreateParserに対応。指定された引数の
// バリデーションをParserクラスに切り出す。Parserはインジェクションで実装を
// 入れ替える。Output処理もParserに統合し外界との変換処理はParserクラスの担当
// として実装する。

type UserCreateHandler interface {
	Handler
	SetUseCase(interface{})
	Create(string) error
}

type userCreateHandler struct {
	interactor usecase.IUserUseCase
	*UserCreateParser
	rhandler
}

//func NewUserCreateHandler(p gateway.IParser) UserCreateHandler {
func NewUserCreateHandler(p *UserCreateParser) UserCreateHandler {
	s := &userCreateHandler{UserCreateParser: p}
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
	format := gateway.NewUserModelFormatParser()
	input := gateway.NewUserInputParser()
	output := gateway.NewUserOutputParser()
	// CreateHandlerに対応する。つまりUserUseCaseの
	// Createに対するInputPortを動的に提供する先を
	// 生成する。
	parser := NewUserCreateParser(format, input, output)
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
			if err := s.Create(data); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "Specify the data in JSON format")
	return cmd
}

func (s *userCreateHandler) Create(data string) error {
	inputPort, ok := interface{}(s).(gateway.IInput)
	if !ok {
		// Inputが存在しない（インターフェースを満たさない）場合は
		// Internal Server Error
		return errors.New(
			codes.InternalServerError,
			"Required Method:[Input()] dose not exist in the specified parser",
		)
	}
	result, err := s.interactor.Create(
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
