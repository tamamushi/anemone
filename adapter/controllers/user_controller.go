/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"
	"github.com/spf13/cobra"

	"anemone/application"
	"anemone/application/usecase"
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
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("data")
			id, _ := cmd.Flags().GetString("id")

			if len(args) > 0 {
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
			} else {
				fmt.Println("no target subcommand")
			}
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
