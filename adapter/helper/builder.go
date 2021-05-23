/* vim: set ts=4 sw=4: */

package helper

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

var instance = map[string]Builder{}
var once sync.Once

func GetBuilderInstance(s string) Builder {
	once.Do(func() {
		instance[s] = &builder{}
	})
	return instance[s]
}

func (s *builder) AddCommand(cmd *cobra.Command) {
	s.commands = append(s.commands, cmd)
}

func (s *builder) GetCommands() []*cobra.Command {
	return s.commands
}
