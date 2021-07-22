/* vim: set ts=4 sw=4: */

package presenter

import (
//"anemone/adapter/gateway"
)

type cliPresenter struct {
}

func NewCLIPresenter() *cliPresenter {
	return &cliPresenter{}
}

func (p *cliPresenter) Command() []string {
	//return []string{"user", "create"}
	return []string{"user", "update"}
	//return []string{"user", "findbyid"}
	//return []string{"user"}
}

func (p *cliPresenter) Arguments() []string {
	return []string{
		"--data",
		"{ \"id\": \"xxxx-xxxx-xxxx-xxx\", \"email\": \"t.koube.cp@gmail.com\" }",
		//"--id",
		//"1000-0001-00AA-BBBB",
	}
}
