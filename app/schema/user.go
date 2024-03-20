package schema

type SignUpRequestBody struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type GetUserRequestQueryParams struct {
	UserId string `json:"UserId"`
}

type PreSignUpQueryExpressionAttributeValueInput struct {
	Email string `json:":email"`
}

type PostConfirmationUpdateItemKeyInput struct {
	UserId string `json:"UserId"`
}

type PostConfirmUpdateItemExpressionAttributeValueInput struct {
	Status string `json:":status"`
}

type SignUpResponse struct {
	UserId                  string `json:"UserId"`
	CodeDeliveryDestination string `json:"CodeDeliveryDestination"`
	CodeDeliveryChannel     string `json:"CodeDeliveryChannel"`
}
