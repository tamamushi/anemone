/* vim: set ts=4 sw=4: */

package application

import (
	"sync"

	"github.com/spf13/cobra"
)

type Builder struct {
	commands []*cobra.Command
}

var instance *Builder
var once sync.Once

func GetBuilderInstance() *Builder {
	once.Do(func() {
		instance = &Builder{}
	})
	return instance
}

func (s *Builder) AddCommand(cmd *cobra.Command) {
	s.commands = append(s.commands, cmd)
}

func (s *Builder) GetCommands() []*cobra.Command {
	return s.commands
}
