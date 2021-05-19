/* vim: set ts=4 sw=4: */

package helper

type ArgumentBuilder interface {
	AddArgs(p string, v string)
	AddCommand(str ...string)
	GetArgString() []string
}

type argumentBuilder struct {
	command []string
	args    []string
}

func NewArgumentBuilder() ArgumentBuilder {
	return &argumentBuilder{}
}

func (c *argumentBuilder) AddArgs(p string, v string) {
	c.args = append(c.args, p, v)
}

func (c *argumentBuilder) AddCommand(str ...string) {
	c.command = append(c.command, str...)
}

func (c *argumentBuilder) GetArgString() []string {
	return append(c.command, c.args...)
}
