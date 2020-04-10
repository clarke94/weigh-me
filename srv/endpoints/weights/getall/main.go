package main

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/clarke94/weigh-me/srv/internal/driver"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	id := request.PathParameters["id"]
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return Response{StatusCode: 404, Body: "UserID required"}, err
	}

	// call the GetAllWeights function ti retrieve all weights
	weights, err := driver.GetAllWeightsByUserID(intID)
	if err != nil {
		return Response{StatusCode: 404, Body: err.Error()}, err
	}

	body, err := json.Marshal(weights)
	if err != nil {
		return Response{StatusCode: 404, Body: err.Error()}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
