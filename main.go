package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"image/png"
	"math"
	"strconv"
)

func sanitizeInitials(initials string) string {
	if len(initials) == 0 {
		return "A"
	}

	if len(initials) > 2 {
		return initials[:2]
	}

	return initials
}

func sanitizeSize(size string) (int, error) {
	if len(size) == 0 {
		return 40, nil
	}

	floatSize, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return 0, err
	}
	floatSize = math.Min(floatSize, 150)
	floatSize = math.Max(floatSize, 10)

	return int(floatSize), nil
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	encoder := png.Encoder{png.BestSpeed, nil}
	buf := new(bytes.Buffer)

	initials := sanitizeInitials(request.QueryStringParameters["initials"])
	size, err := sanitizeSize(request.QueryStringParameters["size"])

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       err.Error(),
		}, nil
	}

	img, err := generateImage(initials, size)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       err.Error(),
		}, nil
	}

	encoder.Encode(buf, img)

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "image/png"},
		Body:            base64.StdEncoding.EncodeToString(buf.Bytes()),
		IsBase64Encoded: true,
	}, nil

}

func main() {
	lambda.Start(HandleRequest)
}
