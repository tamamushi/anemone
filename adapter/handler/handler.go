/* vim: set ts=4 sw=4: */

package handler

import (
	"encoding/json"
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	"anemone/model"
)

type Handler interface {
	Handle() *cobra.Command
	GetHandle() Handler
	SetHandle(func() *cobra.Command)
	SetGateway(*gateway.Gateway)
	GetGateway() *gateway.Gateway
	AddSetter(string, func(interface{}))
	GetSetters() map[string]func(interface{})
	Input(string) interface{}
}

type rhandler struct {
	handle     *func() *cobra.Command
	gateway    *gateway.Gateway
	interactor map[string]func(interface{})
}

func (h *rhandler) GetHandle() Handler {
	return h
}

func (h *rhandler) Handle() *cobra.Command {
	return (*h.handle)()
}

func (h *rhandler) SetHandle(f func() *cobra.Command) {
	h.handle = &f
}

func (h *rhandler) SetGateway(g *gateway.Gateway) {
	h.gateway = g
}

func (h *rhandler) GetGateway() *gateway.Gateway {
	return h.gateway
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

func (h *rhandler) Input(data string) interface{} {

	var p model.User
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		panic(err)
	}
	return p
}
