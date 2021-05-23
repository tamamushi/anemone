/* vim: set ts=4 sw=4: */

package test

import (
	"github.com/spf13/cobra"

	"anemone/adapter/helper"
)

type TestArgument struct {
	Name  string
	Value string
}

func SetupRootCMD(sub string, a *TestArgument) (*cobra.Command, helper.ArgumentBuilder) {

	cmd := &cobra.Command{
		Use:   "anemone",
		Short: "A brief description of your application",

		// Usageは出さない
		SilenceUsage: true,
	}
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	args := helper.NewArgumentBuilder()
	if len(sub) != 0 {
		args.AddCommand(sub)
	}
	if a != nil {
		args.AddArgs(a.Name, a.Value)
	}
	return cmd, args
}
