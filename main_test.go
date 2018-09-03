package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestSanitizeInitials(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "A"},
		{"A", "A"},
		{"AA", "AA"},
		{"AAA", "AA"},
	}

	for _, test := range tests {
		output := sanitizeInitials(test.input)
		if output != test.expected {
			t.Errorf("sanitizeInitials(%s): expected %s, got %s", test.input, test.expected, output)
		}
	}
}

func TestSanitizeSize(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		expectedErr bool
	}{
		{"", 40, false},
		{"9", 10, false},
		{"40", 40, false},
		{"151", 150, false},
		{"aaa", 0, true},
	}

	for _, test := range tests {
		output, outputErr := sanitizeSize(test.input)
		if output != test.expected {
			t.Errorf(
				"sanitizeSize(%s): expected %d, got %d",
				test.input, test.expected, output,
			)
		}
		if (outputErr != nil) != test.expectedErr {
			t.Errorf(
				"sanitizeSize(%s): expected err: %t, got err: %t",
				test.input, test.expectedErr, outputErr != nil,
			)
		}
	}
}

func TestHandleRequest(t *testing.T) {
	params := map[string]string{"initials": "AB", "size": "50"}
	req := events.APIGatewayProxyRequest{QueryStringParameters: params}
	resp1, _ := HandleRequest(req)
	resp2, _ := HandleRequest(req)

	if resp1.Body == resp2.Body {
		t.Error("Expected response body (image content) to differ")
	}
}
