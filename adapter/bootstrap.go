/* vim: set ts=4 sw=4: */

package adapter

import (
	"github.com/spf13/cobra"

	_ "anemone/adapter/controllers"
	"anemone/application"
)

type Response struct {
	Err error
	Cmd *cobra.Command
}

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "anemone",
		Short: "A brief description of your application",
	}
	cobra.OnInitialize(initConfig)

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

func Bootstrap() {

	cmd := NewCmdRoot()
	blder := application.GetBuilderInstance()
	cmd.AddCommand(blder.GetCommands()...)
	cmd.SetArgs([]string{"user"})
	cobra.CheckErr(cmd.Execute())
}

func init() {
}

func initConfig() {
}
