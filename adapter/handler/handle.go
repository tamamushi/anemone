/* vim: set ts=4 sw=4: */

package handler

import (
	"github.com/spf13/cobra"
)

type Handler interface {
	GetHandle() Handler
	SetHandle(func() *cobra.Command)
	AddSetter(string, func(interface{}))
	GetSetters() map[string]func(interface{})
}

type rhandler struct {
	Handle     *func() *cobra.Command
	interactor map[string]func(interface{})
}

func (h *rhandler) GetHandle() Handler {
	return h
}

func (h *rhandler) SetHandle(f func() *cobra.Command) {
	h.Handle = &f
}

func (h *rhandler) AddSetter(k string, f func(interface{})) {
	if h.interactor == nil {
		i := map[string]func(interface{}){}
		h.interactor = i
	}
	h.interactor[k] = f
}

func (h *rhandler) GetSetters() map[string]func(interface{}) {
	return h.interactor
}
