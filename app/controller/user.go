package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	gorillaSchema "github.com/gorilla/schema"
	"github.com/serverless-boilerplate/serverless-chi/app/helper"
	"github.com/serverless-boilerplate/serverless-chi/app/model"
	"github.com/serverless-boilerplate/serverless-chi/app/schema"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	body := schema.SignUpRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}
	clientId := os.Getenv("COGNITO_USERPOOL_CLIENT_ID")
	signUpResponse, err := helper.CognitoIdpClient.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(clientId),
		Username: aws.String(body.Username),
		Password: aws.String(body.Password),
		UserAttributes: []types.AttributeType{{
			Name:  aws.String("email"),
			Value: aws.String(body.Email),
		}},
	})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	userId := *signUpResponse.UserSub
	now := time.Now().Unix()
	user := model.UserModel{
		UserId:    userId,
		Username:  body.Username,
		Email:     body.Password,
		Status:    "Pending",
		CreatedAt: now,
		UpdatedAt: now,
	}
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	putItemResponse, err := helper.DynamodbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("USER_TABLE")),
		Item:      item,
	})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	log.Printf("putItemResponse: %v\n", putItemResponse)
	codeDeliveryChannel := string(signUpResponse.CodeDeliveryDetails.DeliveryMedium)
	codeDeliveryDestination := *signUpResponse.CodeDeliveryDetails.Destination

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(schema.SignUpResponse{
		UserId:                  userId,
		CodeDeliveryDestination: codeDeliveryDestination,
		CodeDeliveryChannel:     codeDeliveryChannel,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(schema.CommonResponse{
		Message: http.StatusText(http.StatusNotImplemented),
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(schema.CommonResponse{
		Message: http.StatusText(http.StatusNotImplemented),
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(schema.CommonResponse{
		Message: http.StatusText(http.StatusNotImplemented),
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := schema.GetUserRequestQueryParams{}
	user := model.UserModel{}
	err := gorillaSchema.NewDecoder().Decode(&params, r.URL.Query())
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}
	if helper.IsValidUUID(params.UserId) {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: "Invalid UUID",
		})
		return
	}
	av, err := attributevalue.MarshalMap(params)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	response, err := helper.DynamodbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:       av,
		TableName: aws.String(os.Getenv("USER_TABLE")),
	})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(schema.CommonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	attributevalue.UnmarshalMap(response.Item, &user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(schema.CommonResponse{
		Message: http.StatusText(http.StatusNotImplemented),
	})
}

func DeactivateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(schema.CommonResponse{
		Message: http.StatusText(http.StatusNotImplemented),
	})
}
