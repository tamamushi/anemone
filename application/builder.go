/* vim: set ts=4 sw=4: */

package usecase

import (
	"fmt"
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

var once = sync.Map{}

func GetBuilderInstance(s string) (Builder, error) {
	instance, _ := once.LoadOrStore(s, &builder{})
	if instance, ok := instance.(Builder); ok {
		//fmt.Printf("%s : %p [%s] \n", s, instance, loaded)
		return instance, nil
	}
	msg := "Can't load instance.(%s : %p)\n"
	return nil, fmt.Errorf(msg, s, instance)
}

func (s *builder) AddCommand(cmd *cobra.Command) {
	s.commands = append(s.commands, cmd)
}

func (s *builder) GetCommands() []*cobra.Command {
	return s.commands
}
