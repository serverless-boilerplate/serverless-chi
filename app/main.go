package main

import (
	"os"

	"github.com/serverless-boilerplate/serverless-chi/app/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handlerMap := map[string]interface{}{
		"AppHandler":       handler.AppHandler,
		"AuthorizeHandler": handler.AuthorizeHandler,
	}
	handlerName := os.Getenv("HANDLER_NAME")
	handlerFunc := handlerMap[handlerName]
	lambda.Start(handlerFunc)
}
