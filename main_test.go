package main

import "testing"

func TestSanitize(t *testing.T) {
	params := RequestParams{"AB", 50}
	params.sanitize()

	if params.Initials != "AB" {
		t.Error("Expected initials to be unchanged, got", params.Initials)
	}

	if params.Size != 50 {
		t.Error("Expected size to be unchanged, got", params.Size)
	}

	bad_params := RequestParams{"ABAB", 500}
	bad_params.sanitize()

	if bad_params.Initials != "AB" {
		t.Error("Expected initials to be changed, got", bad_params.Initials)
	}

	if bad_params.Size != 150 {
		t.Error("Expected size to be changed, got", bad_params.Size)
	}
}

func TestHandleRequest(t *testing.T) {
	params := RequestParams{"AB", 50}
	imgString, _ := HandleRequest(params)
	imgString2, _ := HandleRequest(params)

	if imgString == imgString2 {
		t.Error("Expected image strings to differ")
	}

}
