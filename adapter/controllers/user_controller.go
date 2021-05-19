/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"
	"github.com/spf13/cobra"

	"anemone/application"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

type UserController interface {
	Handler() *cobra.Command
	FindById(id string)
	FindAll()
	Create()
	Remove()
	Update()
}

type userController struct {
	interactor usecase.IUserUseCase
}

func NewUserController(u usecase.IUserUseCase) UserController {
	return &userController{u}
}

func (s *userController) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your command",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				// Error コードを返す
				err := errors.New(
					codes.NotEnoughArgument,
					"Required target sub command",
				)
				fmt.Printf("%v", err.Code())
				fmt.Printf("%#v", err)
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			//name, err := cmd.Flags().GetString("data")
			//id, _ := cmd.Flags().GetString("id")

			switch args[0] {
			case "create":
				s.Create()
			case "remove":
				s.Remove()
			case "update":
				s.Update()
			case "findbyid":
				s.FindById("hoge")
			default:
				fmt.Println("no method")
			}
			return nil
		},
	}

	cmd.Flags().String("data", "", "Your name")
	return cmd
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	controller := NewUserController(usecase)
	blder.AddCommand(controller.Handler())
}

func (s *userController) Create() {
	s.interactor.Create()
}

func (s *userController) Remove() {
}

func (s *userController) Update() {
}

func (s *userController) FindById(id string) {
}

func (s *userController) FindAll() {
}
