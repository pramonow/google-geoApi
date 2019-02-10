package main

import (
	"context"
	"encoding/json"
	"gomapservice/geomap"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
// Handler function Using AWS Lambda Proxy Request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ctx := context.Background()

	//required query
	location := request.QueryStringParameters["location"]
	radius := request.QueryStringParameters["radius"]
	name := request.QueryStringParameters["name"]

	//Replace with api key
	key := "API KEY HERE"

	geoParams := map[string]string{
		"location": location,
		"radius":   radius,
		"key":      key,
	}

	//optional query param
	if name != "" {
		geoParams["name"] = name
	}

	//obtains place nearby response to be processed
	googleResp, err := geomap.PlaceNearby(ctx, geoParams)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 400}, err
	}

	jsonString, _ := json.Marshal(googleResp)

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{Body: string(jsonString), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
