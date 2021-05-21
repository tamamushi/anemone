/* vim: set ts=4 sw=4: */

package main

import (
	"anemone/cmd"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(cmd.Handler)
}
