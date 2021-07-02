/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"
	"os"
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

var onceCommands = sync.Map{}

func CommandBuilder(s string) (Builder, error) {
	instance, _ := onceCommands.LoadOrStore(s, &builder{})
	if instance, ok := instance.(Builder); ok {
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

/*
Constructionが失敗した時にmainルーチンへ戻します。
致命的なエラーが発生した場合にエラー処理する必要がある
場合にはFatalConstructionにerrを判定させます。
err時に発生するエラーメッセージを引数で取ります。
*/
func FatalBuilder(err error, msg string) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		panic(msg)
	}
	return true
}
