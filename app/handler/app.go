package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/serverless-boilerplate/serverless-chi/app/route"
)

func AppHandler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	adapter := chiadapter.NewV2(route.App())
	return adapter.ProxyWithContextV2(ctx, req)
}
