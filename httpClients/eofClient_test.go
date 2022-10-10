package httpClients

import (
	"fmt"
	"testing"
	"time"
)

func Test_PostEOF(t *testing.T) {
	client := NewEOFHTTPClient()

	for i := 0; i < 100; i++ {
		go func(i int) {
			err := client.EOFTest("http://127.0.0.1:54872/PostEOF")
			if err != nil {
				fmt.Printf("i:%d test err: %s\n", i, err)
			}
		}(i)
	}
	time.Sleep(time.Second * 2)
}

func Test_UnexpectedEOF(t *testing.T) {
	client := NewEOFHTTPClient()

	for i := 0; i < 100; i++ {
		go func(i int) {
			err := client.EOFTest("http://127.0.0.1:54872/UnexpectedEOF")
			if err != nil {
				fmt.Printf("i:%d test err: %s\n", i, err)
			}
		}(i)
	}
	time.Sleep(time.Second * 2)
}
