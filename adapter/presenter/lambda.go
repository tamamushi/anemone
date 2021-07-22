/* vim: set ts=4 sw=4: */

package presenter

import (
//"anemone/adapter/gateway"
)

type lambdaPresenter struct {
	fieldName string
}

func NewLambdaPresenter(s string) *lambdaPresenter {
	return &lambdaPresenter{s}
}

func (p *lambdaPresenter) SetArgument(m map[string]string) {
	_ = m
}

func (p *lambdaPresenter) Command() []string {
	return []string{"user", "create"}
}

func (p *lambdaPresenter) Arguments() []string {
	return []string{
		"--data",
		"{ \"id\": \"xxxx-xxxx-xxxx-xxx\",\"email\": \"t.koube.cp@gmail.com\"}",
	}
}
