package handler

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/serverless-boilerplate/serverless-chi/app/helper"
	"github.com/serverless-boilerplate/serverless-chi/app/schema"
)

func PostConfirmationHandler(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	key, err := attributevalue.MarshalMap(schema.PostConfirmationUpdateItemKeyInput{
		UserId: event.Request.UserAttributes["sub"],
	})
	if err != nil {
		return event, err
	}
	expressionAttributeValueInput, err := attributevalue.MarshalMap(schema.PostConfirmUpdateItemExpressionAttributeValueInput{
		Status: "Active",
	})
	if err != nil {
		return event, err
	}
	response, err := helper.DynamodbClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(os.Getenv("USER_TABLE")),
		Key:                       key,
		UpdateExpression:          aws.String("#status = :status"),
		ExpressionAttributeNames:  map[string]string{"#status": "Status"},
		ExpressionAttributeValues: expressionAttributeValueInput,
	})
	if err != nil {
		return event, err
	}
	log.Printf("response: %v\n", response)
	return event, nil
}
