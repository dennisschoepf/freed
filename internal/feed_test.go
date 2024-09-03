package feed

import "testing"

func TestAdd(t *testing.T) {
	gotErr := Add("gibberish")

	if gotErr == nil {
		t.Errorf("Expected Add function to throw an error if an invalid URL is provided")
	}
}
