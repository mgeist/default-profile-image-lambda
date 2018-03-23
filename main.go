package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/lambda"
	"image/png"
	"math"
)

type RequestParams struct {
	Initials string `json:"initials"`
	Size     int    `json:"size"`
}

func (params *RequestParams) sanitize() {
	// max length of 2
	if len(params.Initials) > 2 {
		params.Initials = params.Initials[:2]
	}
	// max size of 150
	// min size of 10
	params.Size = int(math.Min(float64(params.Size), 150))
	params.Size = int(math.Max(float64(params.Size), 10))
}

func HandleRequest(params RequestParams) (string, error) {
	encoder := png.Encoder{png.BestSpeed}
	buf := new(bytes.Buffer)

	params.sanitize()

	img, err := generateImage(params.Initials, params.Size)
	if err != nil {
		return "", err
	}

	encoder.Encode(buf, img)
	encodedString := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encodedString, nil
}

func main() {
	lambda.Start(HandleRequest)
}
