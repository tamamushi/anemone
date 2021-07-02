/* vim: set ts=4 sw=4: */

package handler

import (
	"fmt"
	"os"
	"sync"
)

type ConstructorInterface interface {
	Register(Handler)
	Extraction() []Handler
}

type constructor struct {
	handlers []Handler
}

var onceHandler = sync.Map{}

func Constructor(s string) (ConstructorInterface, error) {
	handler, _ := onceHandler.LoadOrStore(s, &constructor{})
	if handler, ok := handler.(ConstructorInterface); ok {
		return handler, nil
	}
	msg := "Can't load handler.(%s : %p)\n"
	return nil, fmt.Errorf(msg, s, handler)
}

func (s *constructor) Register(hdl Handler) {
	s.handlers = append(s.handlers, hdl)
}

func (s *constructor) Extraction() []Handler {
	return s.handlers
}

/*
Constructionが失敗した時にmainルーチンへ戻します。
致命的なエラーが発生した場合にエラー処理する必要がある
場合にはFatalConstructionにerrを判定させます。
err時に発生するエラーメッセージを引数で取ります。
*/
func FatalConstruction(err error, msg string) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		panic(msg)
	}
	return true
}
