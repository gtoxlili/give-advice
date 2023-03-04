package ht

import (
	"fmt"
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {
	url, err := url.Parse("http://127.0.0.1:16503")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(url)
}
