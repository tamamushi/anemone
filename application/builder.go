/* vim: set ts=4 sw=4: */

package application

import (
	"sync"

	"github.com/spf13/cobra"
)

type Builder interface {
	AddCommand(cmd *cobra.Command)
	GetCommands() []*cobra.Command
}

type builder struct {
	commands []*cobra.Command
}

var instance Builder
var once sync.Once

func GetBuilderInstance() Builder {
	once.Do(func() {
		instance = &builder{}
	})
	return instance
}

func (s *builder) AddCommand(cmd *cobra.Command) {
	s.commands = append(s.commands, cmd)
}

func (s *builder) GetCommands() []*cobra.Command {
	return s.commands
}
