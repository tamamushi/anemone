/* vim: set ts=4 sw=4: */
package api

import (
	_ "anemone/api/user"
	"github.com/spf13/cobra"
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

func Handler() Response {
	cmd := NewCmdRoot()
	rootCmd.SetArgs([]string{"test1"})
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}

func initConfig() {
}
