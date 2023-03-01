package ht

import (
	"fmt"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	{
		defer fmt.Println(1111)
	}
	time.Sleep(time.Second)
	defer fmt.Println(2222)
}
