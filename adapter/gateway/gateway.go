/* vim: set ts=4 sw=4: */

/*
Gateway はgatewayである
*/
package gateway

import ()

type Gateway interface {
	Hoge() string
}

type gateway struct {
}

func NewGateway() *gateway {
	return &gateway{}
}

func (g *gateway) Hoge() string {
	return "hoge"
}
