package main

import (
	"fmt"
	"io"
	"net/http"
)

type myLogger struct{}

func (m myLogger) Write(p []byte) (n int, err error) {
	l := len(p)
	fmt.Println(string(p))
	fmt.Printf("Bytes received: %v\n", l)
	return l, nil
}

func main() {
	r, err := http.Get("http://www.kezic.net")
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return
	}
	m := myLogger{}
	io.Copy(m, r.Body)

}
