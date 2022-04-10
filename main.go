package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func jsonRequests(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println(req.QueryStringParameters["ip"])
	ip, success := req.QueryStringParameters["ip"]
	if !success {
		ret := events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}
		return ret, nil
	}

	ret := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "text/plain; charset=utf-8"},
		Body:       fmt.Sprintf("%#v", jsonResponse(ip)),
	}
	return ret, nil

}

func main() {

	lambda.Start(jsonRequests)

}
