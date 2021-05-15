/* vim: set ts=4 sw=4: */
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"mamsi/cmd"
)

func main() {
	lambda.Start(cmd.Handler)
}
