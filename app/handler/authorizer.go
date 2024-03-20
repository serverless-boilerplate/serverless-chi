package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/serverless-boilerplate/serverless-chi/app/helper"
)

var publicRoutes = map[string]bool{
	"GET /api/v1/service/health": true,
	"GET /api/v1/service/time":   true,
	"GET /api/v1/user/signup":    true,
	"GET /api/v1/user/login":     true,
}

func constructAuthorizerResponse(isAuthorized bool, err error) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: isAuthorized,
	}, err
}

func AuthorizeHandler(ctx context.Context, req events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	accessToken := req.Headers["authorization"]
	requestMethod := req.RequestContext.HTTP.Method
	requestPath := req.RequestContext.HTTP.Path
	route := fmt.Sprintf("%s %s", requestMethod, requestPath)
	_, isPublicRoute := publicRoutes[route]
	if isPublicRoute {
		return constructAuthorizerResponse(true, nil)
	}
	result, err := helper.CognitoIdpClient.GetUser(ctx, &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	})
	if err != nil {
		return constructAuthorizerResponse(false, err)
	}
	log.Print(*result.Username)
	return constructAuthorizerResponse(true, nil)
}
