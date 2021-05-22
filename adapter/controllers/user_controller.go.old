/* vim: set ts=4 sw=4: */

package controllers

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	"anemone/application"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// UserController のinterface定義
type UserController interface {
	Handler() *cobra.Command
	Create() error
	Remove(id string) error
	Update() error
	FindById(id string) error
	Finds() error
}

type userController struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserController(u usecase.IUserUseCase, r gateway.IResponse) UserController {
	return &userController{u, r}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	response := gateway.NewResponse()
	controller := NewUserController(usecase, response)
	blder.AddCommand(controller.Handler())
}

func (s *userController) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your command",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				// Error コードを返す
				return errors.New(
					codes.NotEnoughArgument,
					"Required target sub command",
				)
			}
			switch args[0] {
			case "remove", "findbyid":
				id, _ := cmd.Flags().GetString("id")
				if len(id) < 1 {
					return errors.New(
						codes.NotEnoughArgument,
						"Need to specified \x1b[3mID\x1b[0m",
					)
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			//name, err := cmd.Flags().GetString("data")

			switch args[0] {
			case "create":
				return s.Create()
			case "remove":
				id, _ := cmd.Flags().GetString("id")
				err := s.Remove(id)
				if err != nil {
					return err
				}
				s.response.
					cmd.Printf("%s", id)
			case "update":
				return s.Update()
			case "findbyid":
				id, _ := cmd.Flags().GetString("id")
				err := s.FindById(id)
				if err != nil {
					return err
				}
			default:
				return errors.New(
					codes.UnSupportedMethod,
					errors.Messagef("UnSupported called method: %s", args[0]),
				)
			}
			return nil
		},
	}
	cmd.Flags().String("data", "", "Your name")
	cmd.Flags().String("id", "", "Your name")
	return cmd
}

func (s *userController) Create() error {
	s.interactor.Create()
	return nil
}

func (s *userController) Remove(id string) error {

	err := s.interactor.Remove(id)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}

func (s *userController) Update() error {
	return nil
}

func (s *userController) FindById(id string) error {

	_, err := s.interactor.FindById(id)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}

func (s *userController) Finds() error {
	return nil
}
