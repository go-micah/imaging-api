package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-micah/imaging"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	img := imaging.Dist(512, 512)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	b64img := base64.StdEncoding.EncodeToString(buf.Bytes())

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "image/png",
		},
		Body:            b64img,
		IsBase64Encoded: true,
		StatusCode:      http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
