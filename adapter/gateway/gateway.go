/* vim: set ts=4 sw=4: */

package gateway

// このファイルはGatewayのinterfaceを定義する

import (
	"fmt"
	"sync"
)

type IGateway interface {
	SetResponse(interface{})
	Response() string
}

type gateway struct {
	response interface{}
}

var instance *gateway
var once sync.Once

func GetGateway() *gateway {
	once.Do(func() {
		instance = &gateway{}
	})
	return instance
}

func (g *gateway) SetResponse(m interface{}) {
	g.response = m
}

func (g *gateway) Response() string {
	return fmt.Sprintf("%v#", g.response)
}
