package helper

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

var cfg, _ = config.LoadDefaultConfig(context.TODO())
var DynamodbClient = dynamodb.NewFromConfig(cfg)
var CognitoIdpClient = cognitoidentityprovider.NewFromConfig(cfg)

func IsValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
