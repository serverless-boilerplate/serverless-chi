package handler

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/serverless-boilerplate/serverless-chi/app/helper"
	"github.com/serverless-boilerplate/serverless-chi/app/schema"
)

func PreSignUpHandler(event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	expressionAttributeValues, err := attributevalue.MarshalMap(schema.PreSignUpQueryExpressionAttributeValueInput{
		Email: event.Request.UserAttributes["email"],
	})
	if err != nil {
		return event, err
	}
	response, err := helper.DynamodbClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(os.Getenv("USER_TABLE")),
		IndexName:                 aws.String(os.Getenv("USER_EMAIL_INDEX")),
		FilterExpression:          aws.String("#email = :email"),
		ExpressionAttributeNames:  map[string]string{"#email": "Email"},
		ExpressionAttributeValues: expressionAttributeValues,
		Limit:                     aws.Int32(1),
	})
	if err != nil {
		return event, err
	}
	if response.Count > 0 {
		return event, errors.New("email already exist")
	}
	return event, nil
}
