/* vim: set ts=4 sw=4: */

package main

import (
	"anemone/adapter"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(adapter.Bootstrap)
}
