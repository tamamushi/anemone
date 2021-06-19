/* vim: set ts=4 sw=4: */

package adapter

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"

	_ "anemone/adapter/controllers"
	"anemone/adapter/helper"
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
	buffer := &bytes.Buffer{}
	cmd.SetOutput(buffer)

	blder, _ := helper.GetBuilderInstance("root")
	cmd.AddCommand(blder.GetCommands()...)

	args := helper.NewArgumentBuilder()
	args.AddCommand("user", "create")

	jsond := "{ \"id\": \"xxxx-xxxx-xxxx-xxx\", \"email\": \"t.koube.cp@gmail.com\" }"
	cmd.SetArgs([]string{"user", "create", "--data", jsond})
	//fmt.Printf("%s", []string{"user", "create", "--data", jsond})
	cobra.CheckErr(cmd.Execute())
	fmt.Print(buffer)
}

func initConfig() {

}
