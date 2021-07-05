/* vim: set ts=4 sw=4: */

package handler

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
)

type Handler interface {
	Handle(gateway.Gateway) *cobra.Command
	GetHandle() Handler
	SetHandle(func(gateway.Gateway) *cobra.Command)
	AddSetter(string, func(interface{}))
	GetSetters() map[string]func(interface{})
}

type rhandler struct {
	handle     *func(gateway.Gateway) *cobra.Command
	interactor map[string]func(interface{})
}

func (h *rhandler) GetHandle() Handler {
	return h
}

func (h *rhandler) Handle(g gateway.Gateway) *cobra.Command {
	return (*h.handle)(g)
}

func (h *rhandler) SetHandle(f func(gateway.Gateway) *cobra.Command) {
	h.handle = &f
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
