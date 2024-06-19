package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ip := request.RequestContext.Identity.SourceIP
	fmt.Println("Request from IP: ", ip)
	return events.APIGatewayProxyResponse{
		Body:       string(ip),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
