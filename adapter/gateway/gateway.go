/* vim: set ts=4 sw=4: */

/*
Gateway はgatewayである
*/
package gateway

import (
	"fmt"
	"sync"
)

type Gateway struct {
	response interface{}
}

var instance *Gateway
var once sync.Once

func GetGateway() *Gateway {
	once.Do(func() {
		instance = &Gateway{}
	})
	return instance
}

func (g *Gateway) SetResponse(m interface{}) {
	g.response = m
}

func (g *Gateway) TryParse(m interface{}) error {
	fmt.Printf("%v#", m)
	return nil
}

func (g *Gateway) Response() string {
	return fmt.Sprintf("%v#", g.response)
}
